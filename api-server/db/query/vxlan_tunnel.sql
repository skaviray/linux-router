-- name: CreateVxlanTunnel :one
INSERT INTO vxlan_tunnel (
  name,
  tag,
  tunnel_ip,
  local_ip,
  remote_ip,
  remote_mac,
  status
) VALUES (
  $1, $2,$3,$4,$5,$6,$7
)
RETURNING *;

-- name: GetVxlanTunnel :one
SELECT * FROM vxlan_tunnel
WHERE id = $1 LIMIT 1;

-- name: ListVxlanTunnel :many
SELECT * FROM vxlan_tunnel 
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateVxlanTunnel :one
UPDATE vxlan_tunnel 
SET name=$2
WHERE id=$1 
RETURNING *;

-- name: UpdateVxlanStatus :exec
UPDATE vxlan_tunnel 
SET status=$2
WHERE id=$1;
-- RETURNING *;

-- name: DeleteVxlanTunnel :exec
DELETE FROM vxlan_tunnel
WHERE id = $1;