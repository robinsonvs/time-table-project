-- name: GetUserByID :one
SELECT * from users u where u.uuid = $1;

-- name: CreateUser :exec
INSERT INTO users (uuid, name, email, password)
VALUES ($1, $2, $3, $4);

-- name: FindUserByEmail :one
SELECT u.id, u.uuid, u.name, u.email FROM users u WHERE u.email = $1;

-- name: FindUserByID :one
SELECT u.id, u.uuid, u.name, u.email
FROM users u
WHERE u.uuid = $1;

-- name: UpdateUser :exec
UPDATE users SET
                 name = COALESCE(sqlc.narg('name'), name),
                 email = COALESCE(sqlc.narg('email'), email)
WHERE uuid = $1;

-- name: DeleteUser :exec
DELETE FROM users WHERE uuid = $1;

-- name: FindManyUsers :many
SELECT u.id, u.uuid, u.name, u.email
FROM users u
ORDER BY u.name ASC;

-- name: UpdatePassword :exec
UPDATE users SET password = $2 WHERE uuid = $1;

-- name: GetUserPassword :one
SELECT u.password FROM users u WHERE u.uuid = $1;

