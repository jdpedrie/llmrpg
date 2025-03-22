-- name: UpdateInventoryItemActive :one
UPDATE inventory_items
SET active = $2
WHERE id = $1
RETURNING *;

-- name: UpdateInventoryItemDescription :one
UPDATE inventory_items
SET description = $2
WHERE id = $1
RETURNING *;