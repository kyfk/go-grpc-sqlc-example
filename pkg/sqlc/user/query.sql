/* name: Get :one */
SELECT * FROM users
WHERE id = ?;

/* name: Delete :exec */
DELETE FROM users
WHERE id = ?;

/* name: Create :execresult */
INSERT INTO users (id, username, password) VALUES (?, ?, ?);

/* name: Update :exec */
UPDATE users
SET username = ?, password = ?
WHERE id = ?;
