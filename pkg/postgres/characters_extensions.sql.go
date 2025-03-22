// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: characters_extensions.sql

package postgres

import (
	"context"

	"github.com/google/uuid"
)

const GetCharacterAttributesByType = `-- name: GetCharacterAttributesByType :many
SELECT ca.id, ca.name, ca.value, ca.attribute_type, ca.created_at, ca.updated_at FROM character_attributes ca
JOIN character_to_attributes cal ON ca.id = cal.attribute_id
WHERE cal.character_id = $1::uuid AND ca.attribute_type = $2
`

type GetCharacterAttributesByTypeParams struct {
	CharacterID   uuid.UUID `json:"character_id"`
	AttributeType string    `json:"attribute_type"`
}

func (q *Queries) GetCharacterAttributesByType(ctx context.Context, arg GetCharacterAttributesByTypeParams) ([]CharacterAttribute, error) {
	rows, err := q.db.Query(ctx, GetCharacterAttributesByType, arg.CharacterID, arg.AttributeType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []CharacterAttribute{}
	for rows.Next() {
		var i CharacterAttribute
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Value,
			&i.AttributeType,
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

const GetGameCharacters = `-- name: GetGameCharacters :many
SELECT id, name, description, context, active, main_character, game_id, created_at, updated_at
FROM characters
WHERE game_id = $1::uuid
ORDER BY main_character DESC, name ASC
`

func (q *Queries) GetGameCharacters(ctx context.Context, gameID uuid.UUID) ([]Character, error) {
	rows, err := q.db.Query(ctx, GetGameCharacters, gameID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Character{}
	for rows.Next() {
		var i Character
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.Context,
			&i.Active,
			&i.MainCharacter,
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
