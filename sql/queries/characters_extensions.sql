-- name: GetGameCharacters :many
SELECT id, name, description, context, active, main_character, game_id, created_at, updated_at
FROM characters
WHERE game_id = sqlc.arg(game_id)::uuid
ORDER BY main_character DESC, name ASC;

-- name: GetCharacterAttributesByType :many
SELECT ca.* FROM character_attributes ca
JOIN character_to_attributes cal ON ca.id = cal.attribute_id
WHERE cal.character_id = sqlc.arg(character_id)::uuid AND ca.attribute_type = sqlc.arg(attribute_type);
