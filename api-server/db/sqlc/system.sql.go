// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: system.sql

package sqlc

import (
	"context"
)

const getInitialisation = `-- name: GetInitialisation :one
SELECT id, component, initialised FROM system
WHERE component = $1 LIMIT 1
`

func (q *Queries) GetInitialisation(ctx context.Context, component string) (System, error) {
	row := q.db.QueryRowContext(ctx, getInitialisation, component)
	var i System
	err := row.Scan(&i.ID, &i.Component, &i.Initialised)
	return i, err
}

const markInitialisation = `-- name: MarkInitialisation :one
INSERT INTO system (
  component,
  initialised
) VALUES (
  $1,$2
)
RETURNING id, component, initialised
`

type MarkInitialisationParams struct {
	Component   string `json:"component"`
	Initialised bool   `json:"initialised"`
}

func (q *Queries) MarkInitialisation(ctx context.Context, arg MarkInitialisationParams) (System, error) {
	row := q.db.QueryRowContext(ctx, markInitialisation, arg.Component, arg.Initialised)
	var i System
	err := row.Scan(&i.ID, &i.Component, &i.Initialised)
	return i, err
}
