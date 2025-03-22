-- name: UpdateCharacterName :one
UPDATE characters
SET name = $2
WHERE id = $1
RETURNING *;

-- name: UpdateCharacterDescription :one
UPDATE characters
SET description = $2
WHERE id = $1
RETURNING *;

-- name: UpdateCharacterActive :one
UPDATE characters
SET active = $2
WHERE id = $1
RETURNING *;

-- name: UpdateCharacterContext :one
UPDATE characters
SET context = $2
WHERE id = $1
RETURNING *;