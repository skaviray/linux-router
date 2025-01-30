-- name: CreateVlan :one
INSERT INTO vlans (
  name,
  ipaddress,
  netmask,
  lower,
  tag,
  status
) VALUES (
  $1, $2,$3,$4,$5,$6
)
RETURNING *;

-- name: GetVlan :one
SELECT * FROM vlans
WHERE id = $1 LIMIT 1;

-- name: GetVlanByLowerAndTag :many
SELECT * FROM vlans
WHERE tag = $1 AND lower = $2 LIMIT 1;

-- name: ListVlans :many
SELECT * FROM vlans 
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateVlan :one
UPDATE vlans 
SET name=$2
WHERE id=$1 
RETURNING *;

-- name: UpdateStatus :exec
UPDATE vlans 
SET status=$2
WHERE id=$1 
RETURNING *;

-- name: DeleteVlan :exec
DELETE FROM vlans
WHERE id = $1;