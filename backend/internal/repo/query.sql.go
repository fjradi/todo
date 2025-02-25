// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package repo

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const addTodo = `-- name: AddTodo :one
INSERT INTO
    todo (name)
VALUES
    ($1)
RETURNING id, name, is_completed
`

func (q *Queries) AddTodo(ctx context.Context, name string) (Todo, error) {
	row := q.db.QueryRow(ctx, addTodo, name)
	var i Todo
	err := row.Scan(&i.ID, &i.Name, &i.IsCompleted)
	return i, err
}

const getTodos = `-- name: GetTodos :many
SELECT
    id, name, is_completed
FROM
    todo
WHERE
    name ilike $1
ORDER BY
    name
`

func (q *Queries) GetTodos(ctx context.Context, name string) ([]Todo, error) {
	rows, err := q.db.Query(ctx, getTodos, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Todo{}
	for rows.Next() {
		var i Todo
		if err := rows.Scan(&i.ID, &i.Name, &i.IsCompleted); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateTodo = `-- name: UpdateTodo :one
UPDATE
    todo
SET
    name = $1, is_completed = $2
WHERE
    id = $3
RETURNING
    id, name, is_completed
`

type UpdateTodoParams struct {
	Name        string      `json:"name"`
	IsCompleted bool        `json:"is_completed"`
	ID          pgtype.UUID `json:"id"`
}

func (q *Queries) UpdateTodo(ctx context.Context, arg UpdateTodoParams) (Todo, error) {
	row := q.db.QueryRow(ctx, updateTodo, arg.Name, arg.IsCompleted, arg.ID)
	var i Todo
	err := row.Scan(&i.ID, &i.Name, &i.IsCompleted)
	return i, err
}
