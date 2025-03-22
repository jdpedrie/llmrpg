package model

// ActionOutcome represents the possible outcomes of a player action
type ActionOutcome string

const (
	ActionOutcomeSuccess        ActionOutcome = "success"
	ActionOutcomePartialSuccess ActionOutcome = "partial-success"
	ActionOutcomeFailure        ActionOutcome = "failure"
)

// SuccessProbability represents the likelihood of success for a choice
type SuccessProbability string

const (
	SuccessProbabilityNone       SuccessProbability = "none"
	SuccessProbabilityVeryLow    SuccessProbability = "very_low"
	SuccessProbabilityLow        SuccessProbability = "low"
	SuccessProbabilityMedium     SuccessProbability = "medium"
	SuccessProbabilityHigh       SuccessProbability = "high"
	SuccessProbabilityVeryHigh   SuccessProbability = "very_high"
	SuccessProbabilityGuaranteed SuccessProbability = "guaranteed"
)

// LLMAction represents an action taken by the player
type LLMAction struct {
	Choice  string        `json:"choice"`
	Outcome ActionOutcome `json:"outcome"`
}

// LLMHistoryContextEntry represents a semantically retrieved context entry
type LLMHistoryContextEntry struct {
	Question string   `json:"question"`
	Answers  []string `json:"answers"`
}

// LLMRequest represents the request sent to the LLM
type LLMRequest struct {
	Characters     []Character              `json:"characters"`
	Inventory      []InventoryItem          `json:"inventory"`
	Action         LLMAction                `json:"action"`
	History        []LLMAction              `json:"history"`
	Context        []string                 `json:"context"`
	HistoryContext []LLMHistoryContextEntry `json:"history_context,omitempty"`
}

// LLMChoice represents a choice presented to the player
type LLMChoice struct {
	Text        string             `json:"text"`
	Probability SuccessProbability `json:"probability"`
	Context     string             `json:"context"`
}

// LLMCharacterChange represents a change to a character's basic information
type LLMCharacterChange struct {
	ID    string `json:"id"`
	Field string `json:"field"` // "name", "description", "context", "active"
	Value any    `json:"value"` // string, []string, or bool depending on field
}

// LLMAttributeChange represents a change to a character attribute
type LLMAttributeChange struct {
	ID    string `json:"id"`
	Delta int    `json:"delta"`
}

// LLMInventoryChange represents a change to an inventory item
type LLMInventoryChange struct {
	ID          string `json:"id"`
	Active      *bool  `json:"active,omitempty"`
	Description string `json:"description,omitempty"`
}

// LLMResponse represents the response from the LLM
type LLMResponse struct {
	Message          string               `json:"message"`
	Choices          []LLMChoice          `json:"choices"`
	CharacterChanges []LLMCharacterChange `json:"character_changes"`
	AttributeChanges []LLMAttributeChange `json:"attribute_changes"`
	InventoryChanges []LLMInventoryChange `json:"inventory_changes"`
	ContextQueries   []string             `json:"context_queries"`
	Questions        []string             `json:"questions"`
	NewContexts      []string             `json:"new_contexts"`
}
