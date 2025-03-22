-- name: SearchGameContexts :many
-- This query uses vector embeddings for semantic search
-- It finds contexts related to the query based on vector similarity
-- For text queries, we would first generate an embedding from the query text
SELECT gc.*, (gc.embedding <=> $1) as distance
FROM game_contexts gc
WHERE gc.game_id = $2 AND gc.embedding IS NOT NULL
ORDER BY gc.embedding <=> $1
LIMIT $3;