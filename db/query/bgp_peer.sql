-- name: CreateBgpPeer :one
INSERT INTO bgp_peer (
  as_no,
  neighbor_address,
  local_as
) VALUES (
  $1, $2,$3
)
RETURNING *;

-- name: GetBgpPeer :one
SELECT * FROM bgp_peer
WHERE id = $1 LIMIT 1;

-- name: ListBgpPeers :many
SELECT * FROM bgp_peer 
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateBgpPeer :one
UPDATE bgp_peer 
SET name=$2
WHERE id=$1 
RETURNING *;

-- name: DeleteBgpPeer :exec
DELETE FROM bgp_peer
WHERE id = $1;