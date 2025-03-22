-- name: CreateContextQuery :one
INSERT INTO context_queries (
  game_id, query, used
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetContextQueriesForGame :many
SELECT * FROM context_queries
WHERE game_id = $1 AND used = false;

-- name: MarkContextQueriesAsUsed :exec
UPDATE context_queries
SET used = true
WHERE game_id = $1 AND used = false;

-- name: DeleteContextQueries :exec
DELETE FROM context_queries
WHERE game_id = $1;