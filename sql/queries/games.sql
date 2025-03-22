-- name: GetGame :one
SELECT * FROM games
WHERE id = $1 LIMIT 1;

-- name: ListGames :many
SELECT * FROM games
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: ListActiveGames :many
SELECT * FROM games WHERE is_running = true
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: ListGameTemplates :many
SELECT * FROM games
WHERE is_template = true
ORDER BY name ASC;

-- name: CreateGame :one
INSERT INTO games (
  name, description, starting_message, scenario, objectives,
  skills, characteristics, relationship,
  is_template, is_running
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10
)
RETURNING *;

-- name: UpdateGame :one
UPDATE games
SET
  name = $2,
  description = $3,
  starting_message = $4,
  scenario = $5,
  objectives = $6,
  skills = $7,
  characteristics = $8,
  relationship = $9,
  is_template = $10,
  is_running = $11,
  last_activity_time = NOW()
WHERE id = $1
RETURNING *;

-- name: StartGame :one
UPDATE games
SET
  is_running = true,
  is_template = false,
  playthrough_start_time = NOW(),
  last_activity_time = NOW()
WHERE id = $1
RETURNING *;

-- name: EndGame :one
UPDATE games
SET
  is_running = false,
  playthrough_end_time = NOW(),
  last_activity_time = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteGame :exec
DELETE FROM games
WHERE id = $1;
