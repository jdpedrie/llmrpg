-- name: GetHistoryEntry :one
SELECT * FROM history
WHERE id = $1 LIMIT 1;

-- name: ListGameHistory :many
SELECT * FROM history
WHERE game_id = $1
ORDER BY created_at ASC;

-- name: CreateHistoryEntry :one
INSERT INTO history (
  game_id, text, choice, outcome
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: UpdateHistoryEmbedding :exec
UPDATE history
SET embedding = $2
WHERE id = $1;

-- name: SearchSimilarHistory :many
SELECT h.*, (h.embedding <=> $1) as distance
FROM history h
WHERE h.game_id = $2 AND h.embedding IS NOT NULL
ORDER BY h.embedding <=> $1
LIMIT $3;
