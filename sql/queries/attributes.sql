-- name: GetCharacterAttribute :one
SELECT * FROM character_attributes
WHERE id = $1 LIMIT 1;

-- name: ListCharacterAttributesByType :many
SELECT ca.* FROM character_attributes ca
JOIN character_to_attributes cta ON ca.id = cta.attribute_id
WHERE cta.character_id = $1 AND cta.relationship_type = $2
ORDER BY ca.name ASC;

-- name: CreateCharacterAttribute :one
WITH inserted AS (
  INSERT INTO character_attributes (name, value, attribute_type)
  VALUES ($1, $2, sqlc.arg(attribute_type))
  RETURNING *
)
INSERT INTO character_to_attributes (character_id, attribute_id, relationship_type)
SELECT
  sqlc.arg(character_id)::uuid,
  inserted.id,
  sqlc.arg(attribute_type)
FROM inserted
RETURNING *;

-- name: UpdateCharacterAttribute :one
UPDATE character_attributes
SET
  name = $2,
  value = $3,
  attribute_type = $4
WHERE id = $1
RETURNING *;

-- name: DeleteCharacterAttribute :exec
DELETE FROM character_attributes
WHERE id = $1;

-- name: LinkCharacterAttribute :exec
INSERT INTO character_to_attributes (
  character_id, attribute_id, relationship_type
) VALUES (
  $1, $2, $3
);

-- name: UnlinkCharacterAttribute :exec
DELETE FROM character_to_attributes
WHERE character_id = $1 AND attribute_id = $2;
