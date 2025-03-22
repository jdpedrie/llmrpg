package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/jdpedrie/llmrpg/pkg/postgres"
)

// GameContext represents a piece of unstructured context about a game
// that can be retrieved semantically via vector search
type GameContext struct {
	ID        uuid.UUID
	GameID    uuid.UUID
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// FromDBGameContext converts a database game context to a model game context
func FromDBGameContext(gc postgres.GameContext) GameContext {
	return GameContext{
		ID:        gc.ID,
		GameID:    gc.GameID,
		Content:   gc.Content,
		CreatedAt: gc.CreatedAt.Time,
		UpdatedAt: gc.UpdatedAt.Time,
	}
}

// ToGameContext converts a model game context to a database game context
func (gc *GameContext) ToDBParams() postgres.CreateGameContextParams {
	return postgres.CreateGameContextParams{
		GameID:  gc.GameID,
		Content: gc.Content,
	}
}

// HistoryContextEntry represents an entry in a history_context collection
// which pairs questions with their retrieved context answers
type HistoryContextEntry struct {
	Question string
	Answers  []string
}