-- name: CreateUser :one
INSERT INTO users (
  name,
  email,
  hashed_password
)
VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;

-- name: UpdateUser :one
UPDATE users
set hashed_password = $3, name = $2
where id = $1
returning *;

-- name: DeleteUser :exec
delete from users
where id = $1;

