-- name: CreateUser :one
INSERT INTO cz_users (
  role_id,
  first_name,
  middle_name,
  last_name,
  dob,
  country_code,
  phone,
  email,
  salt,
  password,
  status_id,
  create_user
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM cz_users
WHERE id=$1 LIMIT 1;

-- name: Listusers :many
SELECT * FROM cz_users
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateUser :one
UPDATE cz_users
set password = $1
WHERE id = $2
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM cz_users
WHERE id = $1;

-- name: GetUserForUpdate :one
SELECT * FROM cz_users
WHERE id=$1 LIMIT 1
FOR NO KEY UPDATE;