// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: context.sql

package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/pgvector/pgvector-go"
)

const CreateGameContext = `-- name: CreateGameContext :one
INSERT INTO game_contexts (
  game_id, content, embedding
) VALUES (
  $1, $2, $3
)
RETURNING id, game_id, content, embedding, created_at, updated_at
`

type CreateGameContextParams struct {
	GameID    uuid.UUID       `json:"game_id"`
	Content   string          `json:"content"`
	Embedding pgvector.Vector `json:"embedding"`
}

func (q *Queries) CreateGameContext(ctx context.Context, arg CreateGameContextParams) (GameContext, error) {
	row := q.db.QueryRow(ctx, CreateGameContext, arg.GameID, arg.Content, arg.Embedding)
	var i GameContext
	err := row.Scan(
		&i.ID,
		&i.GameID,
		&i.Content,
		&i.Embedding,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const DeleteGameContext = `-- name: DeleteGameContext :exec
DELETE FROM game_contexts
WHERE id = $1
`

func (q *Queries) DeleteGameContext(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, DeleteGameContext, id)
	return err
}

const GetGameContext = `-- name: GetGameContext :one
SELECT id, game_id, content, embedding, created_at, updated_at FROM game_contexts
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetGameContext(ctx context.Context, id uuid.UUID) (GameContext, error) {
	row := q.db.QueryRow(ctx, GetGameContext, id)
	var i GameContext
	err := row.Scan(
		&i.ID,
		&i.GameID,
		&i.Content,
		&i.Embedding,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const ListGameContexts = `-- name: ListGameContexts :many
SELECT id, game_id, content, embedding, created_at, updated_at FROM game_contexts
WHERE game_id = $1
ORDER BY created_at DESC
`

func (q *Queries) ListGameContexts(ctx context.Context, gameID uuid.UUID) ([]GameContext, error) {
	rows, err := q.db.Query(ctx, ListGameContexts, gameID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GameContext{}
	for rows.Next() {
		var i GameContext
		if err := rows.Scan(
			&i.ID,
			&i.GameID,
			&i.Content,
			&i.Embedding,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const SearchSimilarContexts = `-- name: SearchSimilarContexts :many
SELECT gc.id, gc.game_id, gc.content, gc.embedding, gc.created_at, gc.updated_at, (gc.embedding <=> $1) as distance
FROM game_contexts gc
WHERE gc.game_id = $2 AND gc.embedding IS NOT NULL
ORDER BY gc.embedding <=> $1
LIMIT $3
`

type SearchSimilarContextsParams struct {
	Embedding pgvector.Vector `json:"embedding"`
	GameID    uuid.UUID       `json:"game_id"`
	Limit     int32           `json:"limit"`
}

type SearchSimilarContextsRow struct {
	ID        uuid.UUID          `json:"id"`
	GameID    uuid.UUID          `json:"game_id"`
	Content   string             `json:"content"`
	Embedding pgvector.Vector    `json:"embedding"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	UpdatedAt pgtype.Timestamptz `json:"updated_at"`
	Distance  interface{}        `json:"distance"`
}

func (q *Queries) SearchSimilarContexts(ctx context.Context, arg SearchSimilarContextsParams) ([]SearchSimilarContextsRow, error) {
	rows, err := q.db.Query(ctx, SearchSimilarContexts, arg.Embedding, arg.GameID, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []SearchSimilarContextsRow{}
	for rows.Next() {
		var i SearchSimilarContextsRow
		if err := rows.Scan(
			&i.ID,
			&i.GameID,
			&i.Content,
			&i.Embedding,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Distance,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const UpdateGameContextEmbedding = `-- name: UpdateGameContextEmbedding :exec
UPDATE game_contexts
SET embedding = $2
WHERE id = $1
`

type UpdateGameContextEmbeddingParams struct {
	ID        uuid.UUID       `json:"id"`
	Embedding pgvector.Vector `json:"embedding"`
}

func (q *Queries) UpdateGameContextEmbedding(ctx context.Context, arg UpdateGameContextEmbeddingParams) error {
	_, err := q.db.Exec(ctx, UpdateGameContextEmbedding, arg.ID, arg.Embedding)
	return err
}
