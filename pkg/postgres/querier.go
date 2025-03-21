// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Querier interface {
	CreateCharacter(ctx context.Context, arg CreateCharacterParams) (Character, error)
	CreateCharacterAttribute(ctx context.Context, arg CreateCharacterAttributeParams) (CharacterToAttribute, error)
	CreateContextQuery(ctx context.Context, arg CreateContextQueryParams) (ContextQuery, error)
	CreateGame(ctx context.Context, arg CreateGameParams) (Game, error)
	CreateGameContext(ctx context.Context, arg CreateGameContextParams) (GameContext, error)
	CreateHistory(ctx context.Context, arg CreateHistoryParams) (History, error)
	CreateHistoryEntry(ctx context.Context, arg CreateHistoryEntryParams) (History, error)
	CreateInventoryItem(ctx context.Context, arg CreateInventoryItemParams) (InventoryItem, error)
	DeleteCharacter(ctx context.Context, id uuid.UUID) error
	DeleteCharacterAttribute(ctx context.Context, id uuid.UUID) error
	DeleteContextQueries(ctx context.Context, gameID uuid.UUID) error
	DeleteGame(ctx context.Context, id uuid.UUID) error
	DeleteGameContext(ctx context.Context, id uuid.UUID) error
	DeleteInventoryItem(ctx context.Context, id uuid.UUID) error
	EndGame(ctx context.Context, id uuid.UUID) (Game, error)
	GetCharacter(ctx context.Context, id uuid.UUID) (Character, error)
	GetCharacterAttribute(ctx context.Context, id uuid.UUID) (CharacterAttribute, error)
	GetCharacterAttributesByType(ctx context.Context, arg GetCharacterAttributesByTypeParams) ([]CharacterAttribute, error)
	GetContextQueriesForGame(ctx context.Context, gameID uuid.UUID) ([]ContextQuery, error)
	GetGame(ctx context.Context, id uuid.UUID) (Game, error)
	GetGameCharacters(ctx context.Context, gameID uuid.UUID) ([]Character, error)
	GetGameContext(ctx context.Context, id uuid.UUID) (GameContext, error)
	GetGameHistory(ctx context.Context, arg GetGameHistoryParams) ([]History, error)
	GetGameInventory(ctx context.Context, gameID uuid.UUID) ([]InventoryItem, error)
	GetHistoryEntry(ctx context.Context, id uuid.UUID) (History, error)
	GetInventoryItem(ctx context.Context, id uuid.UUID) (InventoryItem, error)
	LinkCharacterAttribute(ctx context.Context, arg LinkCharacterAttributeParams) error
	ListActiveGames(ctx context.Context, arg ListActiveGamesParams) ([]Game, error)
	ListCharacterAttributesByType(ctx context.Context, arg ListCharacterAttributesByTypeParams) ([]CharacterAttribute, error)
	ListCharacters(ctx context.Context, gameID pgtype.UUID) ([]Character, error)
	ListGameContexts(ctx context.Context, gameID uuid.UUID) ([]GameContext, error)
	ListGameHistory(ctx context.Context, gameID uuid.UUID) ([]History, error)
	ListGameInventory(ctx context.Context, gameID uuid.UUID) ([]InventoryItem, error)
	ListGameTemplates(ctx context.Context) ([]Game, error)
	ListGames(ctx context.Context, arg ListGamesParams) ([]Game, error)
	MarkContextQueriesAsUsed(ctx context.Context, gameID uuid.UUID) error
	// This query uses vector embeddings for semantic search
	// It finds contexts related to the query based on vector similarity
	// For text queries, we would first generate an embedding from the query text
	SearchGameContexts(ctx context.Context, arg SearchGameContextsParams) ([]SearchGameContextsRow, error)
	SearchSimilarContexts(ctx context.Context, arg SearchSimilarContextsParams) ([]SearchSimilarContextsRow, error)
	SearchSimilarHistory(ctx context.Context, arg SearchSimilarHistoryParams) ([]SearchSimilarHistoryRow, error)
	StartGame(ctx context.Context, id uuid.UUID) (Game, error)
	UnlinkCharacterAttribute(ctx context.Context, arg UnlinkCharacterAttributeParams) error
	UpdateCharacter(ctx context.Context, arg UpdateCharacterParams) (Character, error)
	UpdateCharacterActive(ctx context.Context, arg UpdateCharacterActiveParams) (Character, error)
	UpdateCharacterAttribute(ctx context.Context, arg UpdateCharacterAttributeParams) (CharacterAttribute, error)
	UpdateCharacterContext(ctx context.Context, arg UpdateCharacterContextParams) (Character, error)
	UpdateCharacterDescription(ctx context.Context, arg UpdateCharacterDescriptionParams) (Character, error)
	UpdateCharacterName(ctx context.Context, arg UpdateCharacterNameParams) (Character, error)
	UpdateGame(ctx context.Context, arg UpdateGameParams) (Game, error)
	UpdateGameContextEmbedding(ctx context.Context, arg UpdateGameContextEmbeddingParams) error
	UpdateHistoryEmbedding(ctx context.Context, arg UpdateHistoryEmbeddingParams) error
	UpdateInventoryItem(ctx context.Context, arg UpdateInventoryItemParams) (InventoryItem, error)
	UpdateInventoryItemActive(ctx context.Context, arg UpdateInventoryItemActiveParams) (InventoryItem, error)
	UpdateInventoryItemDescription(ctx context.Context, arg UpdateInventoryItemDescriptionParams) (InventoryItem, error)
}

var _ Querier = (*Queries)(nil)
