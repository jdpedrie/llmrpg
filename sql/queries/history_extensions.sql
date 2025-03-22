-- name: GetGameHistory :many
SELECT * FROM history
WHERE game_id = $1
ORDER BY created_at DESC
LIMIT $2;

-- name: CreateHistory :one
INSERT INTO history (
  game_id, text, choice, outcome
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;