-- name: CreateVendor :one
INSERT INTO cz_vendors (
  vendor_id,
  vendor_name,
  registration_number,
  website_url,
  contact_number,
  contact_email,
  country_code,
  salt,
  password,
  status_id,
  create_user
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
)
RETURNING *;

-- name: GetVendor :one
SELECT * FROM cz_vendors
WHERE id=$1 LIMIT 1;

-- name: ListVendors :many
SELECT * FROM cz_vendors
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateVendor :one
UPDATE cz_vendors
set password = $1,
password_modifed_at=$2
WHERE id = $3
RETURNING *;

-- name: DeleteVendor :exec
DELETE FROM cz_vendors
WHERE id = $1;

-- name: GetVendorForUpdate :one
SELECT * FROM cz_vendors
WHERE id=$1 LIMIT 1
FOR NO KEY UPDATE;