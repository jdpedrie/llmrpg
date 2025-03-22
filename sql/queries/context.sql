-- name: GetGameContext :one
SELECT * FROM game_contexts
WHERE id = $1 LIMIT 1;

-- name: ListGameContexts :many
SELECT * FROM game_contexts
WHERE game_id = $1
ORDER BY created_at DESC;

-- name: CreateGameContext :one
INSERT INTO game_contexts (
  game_id, content
) VALUES (
  $1, $2
)
RETURNING *;

-- name: UpdateGameContextEmbedding :exec
UPDATE game_contexts
SET embedding = $2
WHERE id = $1;

-- name: SearchSimilarContexts :many
SELECT gc.*, (gc.embedding <=> $1) as distance
FROM game_contexts gc
WHERE gc.game_id = $2 AND gc.embedding IS NOT NULL
ORDER BY gc.embedding <=> $1
LIMIT $3;

-- name: DeleteGameContext :exec
DELETE FROM game_contexts
WHERE id = $1;