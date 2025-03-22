package game

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"text/template"

	"github.com/jdpedrie/llmrpg/model"
	"github.com/pgvector/pgvector-go"
	"github.com/sashabaranov/go-openai"
)

// LLMClient is an interface for communicating with LLM services
type LLMClient interface {
	// StreamCompletions sends a request to the LLM and returns a stream of the response
	StreamCompletions(ctx context.Context, request *model.LLMRequest) (io.Reader, error)

	// GenerateEmbedding generates an embedding vector for the given text
	GenerateEmbedding(ctx context.Context, text string) (pgvector.Vector, error)

	GenerateContent(ctx context.Context, systemPrompt, userPrompt string, streamHandler func(string, bool)) (string, error)
}

// OpenAIClient implements LLMClient using the OpenAI API
type OpenAIClient struct {
	client          *openai.Client
	modelName       string
	embeddingModel  openai.EmbeddingModel
	systemPromptTpl string
	requestSchema   string
	responseSchema  string
}

// NewOpenAIClient creates a new OpenAI client
func NewOpenAIClient(apiKey, modelName string, embeddingModel openai.EmbeddingModel, systemPromptTplPath string) (*OpenAIClient, error) {
	// Read system prompt template
	systemPromptTpl, err := os.ReadFile(systemPromptTplPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read system prompt template: %w", err)
	}

	// Read schema files
	requestSchema, err := os.ReadFile("schema/request.jsonschema")
	if err != nil {
		return nil, fmt.Errorf("failed to read request schema: %w", err)
	}

	responseSchema, err := os.ReadFile("schema/response.jsonschema")
	if err != nil {
		return nil, fmt.Errorf("failed to read response schema: %w", err)
	}

	return &OpenAIClient{
		client:          openai.NewClient(apiKey),
		modelName:       modelName,
		embeddingModel:  embeddingModel,
		systemPromptTpl: string(systemPromptTpl),
		requestSchema:   string(requestSchema),
		responseSchema:  string(responseSchema),
	}, nil
}

// StreamCompletions implements LLMClient.StreamCompletions using the OpenAI API
func (c *OpenAIClient) StreamCompletions(ctx context.Context, request *model.LLMRequest) (io.Reader, error) {
	// Construct the system prompt from the template
	systemPrompt, err := c.constructSystemPrompt(request)
	if err != nil {
		return nil, fmt.Errorf("failed to construct system prompt: %w", err)
	}

	// Prepare user message (the game state)
	userMessage := c.constructUserMessage(request)

	// Create the stream request
	req := openai.ChatCompletionRequest{
		Model: c.modelName,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: systemPrompt,
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: userMessage,
			},
		},
		Stream: true,
	}

	// Get the stream
	stream, err := c.client.CreateChatCompletionStream(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to create chat completion stream: %w", err)
	}

	// Create a pipe to connect the OpenAI stream with our reader
	pr, pw := io.Pipe()

	// Start a goroutine to read from the stream and write to our pipe
	go func() {
		defer pw.Close()
		defer stream.Close()

		for {
			response, err := stream.Recv()
			if err != nil {
				if err == io.EOF {
					return
				}
				// Just log the error and return, which will close the pipe
				fmt.Fprintf(os.Stderr, "Error receiving from stream: %v\n", err)
				return
			}

			// Write the chunk to our pipe
			if response.Choices[0].Delta.Content != "" {
				_, err = pw.Write([]byte(response.Choices[0].Delta.Content))
				if err != nil {
					fmt.Fprintf(os.Stderr, "Error writing to pipe: %v\n", err)
					return
				}
			}
		}
	}()

	return pr, nil
}

// GenerateEmbedding implements LLMClient.GenerateEmbedding using the OpenAI API
func (c *OpenAIClient) GenerateEmbedding(ctx context.Context, text string) (pgvector.Vector, error) {
	// Create the embedding request
	req := openai.EmbeddingRequest{
		Input: []string{text},
		Model: c.embeddingModel,
	}

	// Get the embedding
	resp, err := c.client.CreateEmbeddings(ctx, req)
	if err != nil {
		return pgvector.Vector{}, fmt.Errorf("failed to create embeddings: %w", err)
	}

	if len(resp.Data) == 0 {
		return pgvector.Vector{}, fmt.Errorf("no embeddings returned")
	}

	// Convert the float32 slice to pgvector.Vector
	return pgvector.NewVector(resp.Data[0].Embedding), nil
}

// constructSystemPrompt builds a system prompt from the template
func (c *OpenAIClient) constructSystemPrompt(req *model.LLMRequest) (string, error) {
	// Parse the template
	tmpl, err := template.New("system_prompt").Parse(c.systemPromptTpl)
	if err != nil {
		return "", fmt.Errorf("failed to parse system prompt template: %w", err)
	}

	// Find the main character in the request
	var mainChar *model.Character
	var otherChars []model.Character
	for i, char := range req.Characters {
		if char.MainCharacter {
			mainChar = &req.Characters[i]
		} else {
			otherChars = append(otherChars, char)
		}
	}

	// Get the game background from context
	background := "You are in a thrilling adventure with elements of: " + findScenarioFromContext(req.Context)

	// Data for the template
	data := map[string]interface{}{
		"GamemasterSchema": c.responseSchema,
		"EngineSchema":     c.requestSchema,
		"Background":       background,
		"EndConditions":    "Game ends when the main character completes their objectives or dies.",
		"IsGameStart":      len(req.History) == 0 && req.Action.Choice == "",
		"InitialMessage":   "You find yourself in a new adventure.",
		"MainCharacters":   formatMainCharacter(mainChar, otherChars),
	}

	// Execute the template
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("failed to execute system prompt template: %w", err)
	}

	return buf.String(), nil
}

// constructUserMessage builds a user message with the game state
func (c *OpenAIClient) constructUserMessage(req *model.LLMRequest) string {
	// Serialize the request to JSON for the LLM
	reqJSON, err := json.MarshalIndent(req, "", "  ")
	if err != nil {
		// If serialization fails, provide a fallback message
		return "Process the current game state and provide a response."
	}

	return string(reqJSON)
}

// findScenarioFromContext extracts scenario information from context
func findScenarioFromContext(contexts []string) string {
	for _, ctx := range contexts {
		if strings.HasPrefix(ctx, "Scenario: ") {
			return strings.TrimPrefix(ctx, "Scenario: ")
		}
	}
	return "an adventure in an unknown world"
}

// formatMainCharacter formats the main character information for the system prompt
func formatMainCharacter(main *model.Character, others []model.Character) string {
	if main == nil {
		return "No main character defined."
	}

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("Main Character: %s\n", main.Name))
	sb.WriteString(fmt.Sprintf("Description: %s\n", main.Description))
	sb.WriteString("Skills:\n")
	for _, skill := range main.Skills {
		sb.WriteString(fmt.Sprintf("- %s: %d\n", skill.Name, skill.Value))
	}

	sb.WriteString("\nOther Characters:\n")
	for _, char := range others {
		sb.WriteString(fmt.Sprintf("- %s: %s\n", char.Name, char.Description))
	}

	return sb.String()
}

// MockLLMClient is a mock implementation of LLMClient for testing
type MockLLMClient struct{}

// StreamCompletions returns a predetermined response for testing
func (m *MockLLMClient) StreamCompletions(ctx context.Context, request *model.LLMRequest) (io.Reader, error) {
	// Prepare a mock response
	mockResponse := `You find yourself in a dark cave. Water drips from stalactites above, and the air is damp and cold.
To your left, you see a faint light. To your right, you hear a strange noise.

[END_MESSAGE]
{
  "message": "You find yourself in a dark cave. Water drips from stalactites above, and the air is damp and cold.\nTo your left, you see a faint light. To your right, you hear a strange noise.",
  "choices": [
    {
      "text": "Go towards the light",
      "probability": "high",
      "context": "The light seems safe and inviting."
    },
    {
      "text": "Investigate the noise",
      "probability": "medium",
      "context": "The noise could be a threat, but also an opportunity."
    },
    {
      "text": "Stay where you are and wait",
      "probability": "guaranteed",
      "context": "Staying put is safe but won't advance your quest."
    }
  ],
  "character_changes": [],
  "attribute_changes": [],
  "inventory_changes": [],
  "context_queries": [],
  "questions": [
    "What happened before the character entered the cave?",
    "Does the character have any light sources in their inventory?"
  ],
  "new_contexts": [
    "The hero entered a dark cave with dripping water and cold air.",
    "There is a mysterious light visible to the left side of the cave.",
    "Strange noises can be heard coming from the right side of the cave."
  ]
}`

	return strings.NewReader(mockResponse), nil
}

// GenerateEmbedding returns a mock embedding for testing
func (m *MockLLMClient) GenerateEmbedding(ctx context.Context, text string) (pgvector.Vector, error) {
	// Return a simple mock embedding (a 1536-element vector with all values set to 0.1)
	mockEmbedding := make([]float32, 1536)
	for i := range mockEmbedding {
		mockEmbedding[i] = 0.1
	}
	return pgvector.NewVector(mockEmbedding), nil
}
