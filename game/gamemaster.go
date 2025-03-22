package game

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strings"
	"text/template"

	"github.com/google/uuid"
	"github.com/jdpedrie/llmrpg/pkg/postgres"
)

// GameMaster handles interactions with the LLM and game state updates
type GameMaster struct {
	engine    *Engine
	llmClient LLMClient
	template  *template.Template
}

// NewGameMaster creates a new game master
func NewGameMaster(engine *Engine, llmClient LLMClient) *GameMaster {
	// Load the system prompt template
	tmpl, err := template.ParseFiles("game/system_prompt.tpl")
	if err != nil {
		panic(fmt.Sprintf("failed to load system prompt template: %v", err))
	}

	return &GameMaster{
		engine:    engine,
		llmClient: llmClient,
		template:  tmpl,
	}
}

// ProcessAction processes a player action and updates the game state
func (gm *GameMaster) ProcessAction(
	ctx context.Context,
	qtx *postgres.Queries,
	gameID uuid.UUID,
	choice, outcome string,
	responseChan chan<- string,
	finalChan chan<- bool,
) error {
	defer close(responseChan)
	defer close(finalChan)

	// Get the game state
	dbGame, err := qtx.GetGame(ctx, gameID)
	if err != nil {
		return err
	}

	// Get characters
	characters, err := qtx.GetGameCharacters(ctx, gameID)
	if err != nil {
		return err
	}

	mainCharacter := postgres.Character{}
	nonPlayerCharacters := []postgres.Character{}

	for _, char := range characters {
		if char.MainCharacter {
			mainCharacter = char
		} else {
			nonPlayerCharacters = append(nonPlayerCharacters, char)
		}
	}

	// Get history
	history, err := qtx.GetGameHistory(ctx, postgres.GetGameHistoryParams{
		GameID: gameID,
		Limit:  10,
	})
	if err != nil {
		return err
	}

	// Get semantic context
	semanticContext, err := gm.getSemanticContext(ctx, qtx, gameID, choice)
	if err != nil {
		return err
	}

	// Generate the system prompt
	systemPrompt, err := gm.generateSystemPrompt(dbGame, mainCharacter, nonPlayerCharacters, history, semanticContext)
	if err != nil {
		return err
	}

	// Prepare user prompt
	userPrompt := choice
	if outcome != "" {
		userPrompt = fmt.Sprintf("%s\n\nOutcome: %s", choice, outcome)
	}

	// Call the LLM
	response, err := gm.llmClient.GenerateContent(ctx, systemPrompt, userPrompt, func(token string, isFinal bool) {
		// Split token is a marker for the final updated game state
		if strings.Contains(token, "SPLIT_TOKEN") {
			if isFinal {
				responseChan <- token
				finalChan <- true
			}
			return
		}

		responseChan <- token
		finalChan <- isFinal
	})

	if err != nil {
		return err
	}

	// Save the history entry
	historyParams := postgres.CreateHistoryParams{
		GameID:  gameID,
		Text:    response,
		Choice:  choice,
		Outcome: outcome,
	}

	_, err = qtx.CreateHistory(ctx, historyParams)
	if err != nil {
		return err
	}

	// Mark relevant context queries as used
	return qtx.MarkContextQueriesAsUsed(ctx, gameID)
}

// getSemanticContext retrieves semantically relevant context from the history
func (gm *GameMaster) getSemanticContext(ctx context.Context, qtx *postgres.Queries, gameID uuid.UUID, query string) ([]string, error) {
	// Generate embedding for the query
	embedding, err := gm.llmClient.GenerateEmbedding(ctx, query)
	if err != nil {
		return nil, err
	}

	// Perform similarity search
	contextEntries, err := qtx.SearchSimilarContexts(ctx, postgres.SearchSimilarContextsParams{
		GameID:    gameID,
		Embedding: embedding,
		Limit:     5,
	})
	if err != nil {
		return nil, err
	}

	results := make([]string, 0, len(contextEntries))
	for _, entry := range contextEntries {
		results = append(results, entry.Content)
	}

	return results, nil
}

// generateSystemPrompt creates the system prompt for the LLM
func (gm *GameMaster) generateSystemPrompt(
	game postgres.Game,
	mainCharacter postgres.Character,
	npcs []postgres.Character,
	history []postgres.History,
	semanticContext []string,
) (string, error) {
	// Prepare data for template
	data := map[string]interface{}{
		"Game":            game,
		"MainCharacter":   mainCharacter,
		"NPCs":            npcs,
		"History":         history,
		"SemanticContext": semanticContext,
	}

	// Execute the template
	var buf bytes.Buffer
	if err := gm.template.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}

// JSONSchema represents a JSON schema for LLM responses
type JSONSchema struct {
	Type       string                 `json:"type"`
	Properties map[string]interface{} `json:"properties"`
	Required   []string               `json:"required"`
}

// ParseJSONResponse parses a JSON response from the LLM
func ParseJSONResponse(response string, schema JSONSchema) (map[string]interface{}, error) {
	// Extract JSON content
	jsonStart := strings.Index(response, "{")
	jsonEnd := strings.LastIndex(response, "}")

	if jsonStart == -1 || jsonEnd == -1 || jsonEnd < jsonStart {
		return nil, fmt.Errorf("invalid JSON format in response: %s", response)
	}

	jsonContent := response[jsonStart : jsonEnd+1]

	// Parse JSON
	var result map[string]interface{}
	if err := json.Unmarshal([]byte(jsonContent), &result); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	// Validate required fields
	for _, field := range schema.Required {
		if _, ok := result[field]; !ok {
			return nil, fmt.Errorf("required field '%s' missing in response", field)
		}
	}

	return result, nil
}
