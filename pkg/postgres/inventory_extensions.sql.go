// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: inventory_extensions.sql

package postgres

import (
	"context"

	"github.com/google/uuid"
)

const GetGameInventory = `-- name: GetGameInventory :many
SELECT id, name, description, active, game_id, created_at, updated_at FROM inventory_items
WHERE game_id = $1::uuid
ORDER BY name ASC
`

func (q *Queries) GetGameInventory(ctx context.Context, gameID uuid.UUID) ([]InventoryItem, error) {
	rows, err := q.db.Query(ctx, GetGameInventory, gameID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []InventoryItem{}
	for rows.Next() {
		var i InventoryItem
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.Active,
			&i.GameID,
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
