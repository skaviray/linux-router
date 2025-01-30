// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: vlans.sql

package sqlc

import (
	"context"
)

const createVlan = `-- name: CreateVlan :one
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
RETURNING id, name, ipaddress, netmask, lower, tag, status, created_at
`

type CreateVlanParams struct {
	Name      string `json:"name"`
	Ipaddress string `json:"ipaddress"`
	Netmask   string `json:"netmask"`
	Lower     int64  `json:"lower"`
	Tag       int64  `json:"tag"`
	Status    string `json:"status"`
}

func (q *Queries) CreateVlan(ctx context.Context, arg CreateVlanParams) (Vlan, error) {
	row := q.db.QueryRowContext(ctx, createVlan,
		arg.Name,
		arg.Ipaddress,
		arg.Netmask,
		arg.Lower,
		arg.Tag,
		arg.Status,
	)
	var i Vlan
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Ipaddress,
		&i.Netmask,
		&i.Lower,
		&i.Tag,
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
}

const deleteVlan = `-- name: DeleteVlan :exec
DELETE FROM vlans
WHERE id = $1
`

func (q *Queries) DeleteVlan(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteVlan, id)
	return err
}

const getVlan = `-- name: GetVlan :one
SELECT id, name, ipaddress, netmask, lower, tag, status, created_at FROM vlans
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetVlan(ctx context.Context, id int64) (Vlan, error) {
	row := q.db.QueryRowContext(ctx, getVlan, id)
	var i Vlan
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Ipaddress,
		&i.Netmask,
		&i.Lower,
		&i.Tag,
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
}

const getVlanByLowerAndTag = `-- name: GetVlanByLowerAndTag :many
SELECT id, name, ipaddress, netmask, lower, tag, status, created_at FROM vlans
WHERE tag = $1 AND lower = $2 LIMIT 1
`

type GetVlanByLowerAndTagParams struct {
	Tag   int64 `json:"tag"`
	Lower int64 `json:"lower"`
}

func (q *Queries) GetVlanByLowerAndTag(ctx context.Context, arg GetVlanByLowerAndTagParams) ([]Vlan, error) {
	rows, err := q.db.QueryContext(ctx, getVlanByLowerAndTag, arg.Tag, arg.Lower)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Vlan{}
	for rows.Next() {
		var i Vlan
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Ipaddress,
			&i.Netmask,
			&i.Lower,
			&i.Tag,
			&i.Status,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listVlans = `-- name: ListVlans :many
SELECT id, name, ipaddress, netmask, lower, tag, status, created_at FROM vlans 
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListVlansParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListVlans(ctx context.Context, arg ListVlansParams) ([]Vlan, error) {
	rows, err := q.db.QueryContext(ctx, listVlans, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Vlan{}
	for rows.Next() {
		var i Vlan
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Ipaddress,
			&i.Netmask,
			&i.Lower,
			&i.Tag,
			&i.Status,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateStatus = `-- name: UpdateStatus :exec
UPDATE vlans 
SET status=$2
WHERE id=$1 
RETURNING id, name, ipaddress, netmask, lower, tag, status, created_at
`

type UpdateStatusParams struct {
	ID     int64  `json:"id"`
	Status string `json:"status"`
}

func (q *Queries) UpdateStatus(ctx context.Context, arg UpdateStatusParams) error {
	_, err := q.db.ExecContext(ctx, updateStatus, arg.ID, arg.Status)
	return err
}

const updateVlan = `-- name: UpdateVlan :one
UPDATE vlans 
SET name=$2
WHERE id=$1 
RETURNING id, name, ipaddress, netmask, lower, tag, status, created_at
`

type UpdateVlanParams struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func (q *Queries) UpdateVlan(ctx context.Context, arg UpdateVlanParams) (Vlan, error) {
	row := q.db.QueryRowContext(ctx, updateVlan, arg.ID, arg.Name)
	var i Vlan
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Ipaddress,
		&i.Netmask,
		&i.Lower,
		&i.Tag,
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
}
