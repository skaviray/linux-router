-- name: CreateInterface :one
INSERT INTO interfaces (
  macaddress,
  ipaddress,
  mtu,
  name
) VALUES (
  $1, $2,$3,$4
)
RETURNING *;

-- name: GetInterface :one
SELECT * FROM interfaces
WHERE id = $1 LIMIT 1;

-- name: ListInterfaces :many
SELECT * FROM interfaces 
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateInterface :one
UPDATE interfaces 
SET name=$2
WHERE id=$1 
RETURNING *;

-- name: DeleteInterface :exec
DELETE FROM interfaces
WHERE id = $1;