-- name: GetCharacter :one
SELECT * FROM characters
WHERE id = $1 LIMIT 1;

-- name: ListCharacters :many
SELECT * FROM characters
WHERE game_id = $1
ORDER BY name ASC;

-- name: CreateCharacter :one
INSERT INTO characters (
  name, description, context, active, main_character, game_id
) VALUES (
  $1, $2, $3, $4, $5, sqlc.arg(game_id)::uuid
)
RETURNING *;

-- name: UpdateCharacter :one
UPDATE characters
SET
  name = $2,
  description = $3,
  context = $4,
  active = $5,
  main_character = $6
WHERE id = $1
RETURNING *;

-- name: DeleteCharacter :exec
DELETE FROM characters
WHERE id = $1;
