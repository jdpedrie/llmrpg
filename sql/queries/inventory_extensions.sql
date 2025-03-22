-- name: GetGameInventory :many
SELECT id, name, description, active, game_id, created_at, updated_at FROM inventory_items
WHERE game_id = sqlc.arg(game_id)::uuid
ORDER BY name ASC;
