-- name: CreateCustomer :one
INSERT INTO customers (
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

-- name: GetCustomer :one
SELECT * FROM customers
WHERE id=$1 LIMIT 1;

-- name: ListCustomers :many
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

-- name: GetCustomerForUpdate :one
SELECT * FROM customers
WHERE id=$1 LIMIT 1
FOR NO KEY UPDATE;