-- name: CreateCustomer :one
INSERT INTO customers (
  role_id,
  first_name,
  last_name,
  dob,
  mobile,
  email,
  password,
  status
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8
)
RETURNING *;

-- name: GetCustomer :one
SELECT * FROM customers
WHERE id=$1 LIMIT 1;

-- name: ListAuthors :many
SELECT * FROM customers
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateCustomer :one
UPDATE customers
set password = $1
WHERE id = $2
RETURNING *;

-- name: DeleteCustomer :exec
DELETE FROM customers
WHERE id = $1;