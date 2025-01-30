-- name: CreateBgpAdvertisement :one
INSERT INTO bgp_advertisement (
  name,
  destination_cidr
) VALUES (
  $1,$2
)
RETURNING *;

-- name: GetBgpAdvertisement :one
SELECT * FROM bgp_advertisement
WHERE id = $1 LIMIT 1;

-- name: ListBgpAdvertisements :many
SELECT * FROM bgp_advertisement 
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateBgpAdvertisement :one
UPDATE bgp_advertisement 
SET name=$2
WHERE id=$1 
RETURNING *;

-- name: DeleteBgpAdvertisement :exec
DELETE FROM bgp_advertisement
WHERE id = $1;