-- name: GetTodos :many
SELECT
    id, name, is_completed
FROM
    todo
ORDER BY
    name;

-- name: AddTodo :one
INSERT INTO
    todo (name)
VALUES
    ($1)
RETURNING id, name, is_completed;

-- name: UpdateTodo :one
UPDATE
    todo
SET
    name = $1, is_completed = $2
WHERE
    id = $3
RETURNING
    id, name, is_completed;