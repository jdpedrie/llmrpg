-- name: GetInventoryItem :one
SELECT * FROM inventory_items
WHERE id = $1 LIMIT 1;

-- name: ListGameInventory :many
SELECT * FROM inventory_items
WHERE game_id = $1
ORDER BY name ASC;

-- name: CreateInventoryItem :one
INSERT INTO inventory_items (
  name, description, active, game_id
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: UpdateInventoryItem :one
UPDATE inventory_items
SET 
  name = $2,
  description = $3,
  active = $4
WHERE id = $1
RETURNING *;

-- name: DeleteInventoryItem :exec
DELETE FROM inventory_items
WHERE id = $1;
