// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package postgres

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/pgvector/pgvector-go"
)

type Character struct {
	ID            uuid.UUID          `json:"id"`
	Name          string             `json:"name"`
	Description   pgtype.Text        `json:"description"`
	Context       []string           `json:"context"`
	Active        bool               `json:"active"`
	MainCharacter bool               `json:"main_character"`
	GameID        pgtype.UUID        `json:"game_id"`
	CreatedAt     pgtype.Timestamptz `json:"created_at"`
	UpdatedAt     pgtype.Timestamptz `json:"updated_at"`
}

type CharacterAttribute struct {
	ID            uuid.UUID          `json:"id"`
	Name          string             `json:"name"`
	Value         int16              `json:"value"`
	AttributeType string             `json:"attribute_type"`
	CreatedAt     pgtype.Timestamptz `json:"created_at"`
	UpdatedAt     pgtype.Timestamptz `json:"updated_at"`
}

type CharacterToAttribute struct {
	CharacterID      uuid.UUID          `json:"character_id"`
	AttributeID      uuid.UUID          `json:"attribute_id"`
	RelationshipType string             `json:"relationship_type"`
	CreatedAt        pgtype.Timestamptz `json:"created_at"`
}

type Game struct {
	ID                   uuid.UUID          `json:"id"`
	Name                 string             `json:"name"`
	Description          pgtype.Text        `json:"description"`
	StartingMessage      pgtype.Text        `json:"starting_message"`
	Scenario             pgtype.Text        `json:"scenario"`
	Objectives           pgtype.Text        `json:"objectives"`
	Skills               []string           `json:"skills"`
	Characteristics      []string           `json:"characteristics"`
	Relationship         []string           `json:"relationship"`
	IsTemplate           bool               `json:"is_template"`
	IsRunning            bool               `json:"is_running"`
	PlaythroughStartTime pgtype.Timestamptz `json:"playthrough_start_time"`
	PlaythroughEndTime   pgtype.Timestamptz `json:"playthrough_end_time"`
	LastActivityTime     pgtype.Timestamptz `json:"last_activity_time"`
	CreatedAt            pgtype.Timestamptz `json:"created_at"`
	UpdatedAt            pgtype.Timestamptz `json:"updated_at"`
}

type GameContext struct {
	ID        uuid.UUID          `json:"id"`
	GameID    uuid.UUID          `json:"game_id"`
	Content   string             `json:"content"`
	Embedding pgvector.Vector    `json:"embedding"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	UpdatedAt pgtype.Timestamptz `json:"updated_at"`
}

type History struct {
	ID        uuid.UUID          `json:"id"`
	GameID    uuid.UUID          `json:"game_id"`
	Text      string             `json:"text"`
	Choice    string             `json:"choice"`
	Outcome   string             `json:"outcome"`
	Embedding pgvector.Vector    `json:"embedding"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
}

type InventoryItem struct {
	ID          uuid.UUID          `json:"id"`
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Active      bool               `json:"active"`
	GameID      uuid.UUID          `json:"game_id"`
	CreatedAt   pgtype.Timestamptz `json:"created_at"`
	UpdatedAt   pgtype.Timestamptz `json:"updated_at"`
}
