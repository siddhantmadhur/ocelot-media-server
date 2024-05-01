// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package storage

import (
	"context"
	"database/sql"
)

const getProfiles = `-- name: GetProfiles :many
SELECT id, name FROM profiles
`

type GetProfilesRow struct {
	ID   int64
	Name sql.NullString
}

func (q *Queries) GetProfiles(ctx context.Context) ([]GetProfilesRow, error) {
	rows, err := q.db.QueryContext(ctx, getProfiles)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetProfilesRow
	for rows.Next() {
		var i GetProfilesRow
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
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