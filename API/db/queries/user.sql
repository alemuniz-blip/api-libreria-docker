-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = ? LIMIT 1;

-- name: GetAllUsers :many
SELECT * FROM users;

-- name: GetUserById :one
SELECT * FROM users
WHERE id = ?;

-- name: UpdateUser :exec
UPDATE users
SET name = ?, email = ?
WHERE id = ?;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = ?;