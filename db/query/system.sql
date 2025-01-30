-- name: MarkInitialisation :one
INSERT INTO system (
  component,
  initialised
) VALUES (
  $1,$2
)
RETURNING *;

-- name: GetInitialisation :one
SELECT * FROM system
WHERE component = $1 LIMIT 1;