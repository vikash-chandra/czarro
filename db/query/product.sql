-- name: CreateProduct :one
INSERT INTO cz_products (
  title,
  short_name,
  description,
  sms_noti,
  email_noti,
  call_noti,
  image,
  currency_id,
  price,
  status_id,
  create_user
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
)
RETURNING *;

-- name: GetProduct :one
SELECT * FROM cz_products
WHERE id=$1 LIMIT 1;

-- name: ListProducts :many
SELECT * FROM cz_products
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateProducts :one
UPDATE cz_products
SET 
    title = $1,
    short_name = $2,
    description = $3,
    sms_noti = $4,
    email_noti = $5,
    call_noti = $6,
    image = $7,
    currency_id = $8,
    price = $9,
    status_id = $10,
    modify_user = $11,
    modified_at = $12
WHERE id = $13
RETURNING *;

-- name: DeleteProduct :exec
DELETE FROM cz_products
WHERE id = $1;

-- name: GetProductForUpdate :one
SELECT * FROM cz_products
WHERE id=$1 LIMIT 1
FOR NO KEY UPDATE;