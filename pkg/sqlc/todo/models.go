// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package todo

import (
	"database/sql"
	"time"
)

type Todo struct {
	ID        string
	UserID    string
	Content   string
	DueTo     sql.NullTime
	CreatedAt time.Time
	UpdatedAt time.Time
}
