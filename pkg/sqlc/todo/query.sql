/* name: Get :one */
SELECT * FROM todos
WHERE id = ?;

/* name: GetByUser :many */
SELECT * FROM todos
WHERE user_id = ?;

/* name: Delete :exec */
DELETE FROM todos
WHERE id = ?;

/* name: Create :execresult */
INSERT INTO todos (id, user_id, content, due_to) VALUES (?, ?, ?, ?);

/* name: Update :exec */
UPDATE todos
SET content = ?, due_to = ?
WHERE id = ?;
