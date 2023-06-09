// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: query.sql

package todo

import (
	"context"
	"database/sql"
)

const create = `-- name: Create :execresult
INSERT INTO todos (id, user_id, content, due_to) VALUES (?, ?, ?, ?)
`

type CreateParams struct {
	ID      string
	UserID  string
	Content string
	DueTo   sql.NullTime
}

func (q *Queries) Create(ctx context.Context, arg CreateParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, create,
		arg.ID,
		arg.UserID,
		arg.Content,
		arg.DueTo,
	)
}

const delete = `-- name: Delete :exec
DELETE FROM todos
WHERE id = ?
`

func (q *Queries) Delete(ctx context.Context, id string) error {
	_, err := q.db.ExecContext(ctx, delete, id)
	return err
}

const get = `-- name: Get :one
SELECT id, user_id, content, due_to, created_at, updated_at FROM todos
WHERE id = ?
`

func (q *Queries) Get(ctx context.Context, id string) (Todo, error) {
	row := q.db.QueryRowContext(ctx, get, id)
	var i Todo
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Content,
		&i.DueTo,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getByUser = `-- name: GetByUser :many
SELECT id, user_id, content, due_to, created_at, updated_at FROM todos
WHERE user_id = ?
`

func (q *Queries) GetByUser(ctx context.Context, userID string) ([]Todo, error) {
	rows, err := q.db.QueryContext(ctx, getByUser, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Todo
	for rows.Next() {
		var i Todo
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Content,
			&i.DueTo,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const update = `-- name: Update :exec
UPDATE todos
SET content = ?, due_to = ?
WHERE id = ?
`

type UpdateParams struct {
	Content string
	DueTo   sql.NullTime
	ID      string
}

func (q *Queries) Update(ctx context.Context, arg UpdateParams) error {
	_, err := q.db.ExecContext(ctx, update, arg.Content, arg.DueTo, arg.ID)
	return err
}
