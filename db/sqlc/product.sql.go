// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: product.sql

package db

import (
	"context"
	"time"
)

const createProduct = `-- name: CreateProduct :one
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
RETURNING id, title, short_name, description, sms_noti, email_noti, call_noti, image, currency_id, price, status_id, create_user, modify_user, created_at, modified_at
`

type CreateProductParams struct {
	Title       string  `json:"title"`
	ShortName   string  `json:"short_name"`
	Description string  `json:"description"`
	SmsNoti     bool    `json:"sms_noti"`
	EmailNoti   bool    `json:"email_noti"`
	CallNoti    bool    `json:"call_noti"`
	Image       string  `json:"image"`
	CurrencyID  int32   `json:"currency_id"`
	Price       float64 `json:"price"`
	StatusID    int32   `json:"status_id"`
	CreateUser  int64   `json:"create_user"`
}

func (q *Queries) CreateProduct(ctx context.Context, arg CreateProductParams) (CzProduct, error) {
	row := q.db.QueryRow(ctx, createProduct,
		arg.Title,
		arg.ShortName,
		arg.Description,
		arg.SmsNoti,
		arg.EmailNoti,
		arg.CallNoti,
		arg.Image,
		arg.CurrencyID,
		arg.Price,
		arg.StatusID,
		arg.CreateUser,
	)
	var i CzProduct
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.ShortName,
		&i.Description,
		&i.SmsNoti,
		&i.EmailNoti,
		&i.CallNoti,
		&i.Image,
		&i.CurrencyID,
		&i.Price,
		&i.StatusID,
		&i.CreateUser,
		&i.ModifyUser,
		&i.CreatedAt,
		&i.ModifiedAt,
	)
	return i, err
}

const deleteProduct = `-- name: DeleteProduct :exec
DELETE FROM cz_products
WHERE id = $1
`

func (q *Queries) DeleteProduct(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteProduct, id)
	return err
}

const getProduct = `-- name: GetProduct :one
SELECT id, title, short_name, description, sms_noti, email_noti, call_noti, image, currency_id, price, status_id, create_user, modify_user, created_at, modified_at FROM cz_products
WHERE id=$1 LIMIT 1
`

func (q *Queries) GetProduct(ctx context.Context, id int32) (CzProduct, error) {
	row := q.db.QueryRow(ctx, getProduct, id)
	var i CzProduct
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.ShortName,
		&i.Description,
		&i.SmsNoti,
		&i.EmailNoti,
		&i.CallNoti,
		&i.Image,
		&i.CurrencyID,
		&i.Price,
		&i.StatusID,
		&i.CreateUser,
		&i.ModifyUser,
		&i.CreatedAt,
		&i.ModifiedAt,
	)
	return i, err
}

const getProductForUpdate = `-- name: GetProductForUpdate :one
SELECT id, title, short_name, description, sms_noti, email_noti, call_noti, image, currency_id, price, status_id, create_user, modify_user, created_at, modified_at FROM cz_products
WHERE id=$1 LIMIT 1
FOR NO KEY UPDATE
`

func (q *Queries) GetProductForUpdate(ctx context.Context, id int32) (CzProduct, error) {
	row := q.db.QueryRow(ctx, getProductForUpdate, id)
	var i CzProduct
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.ShortName,
		&i.Description,
		&i.SmsNoti,
		&i.EmailNoti,
		&i.CallNoti,
		&i.Image,
		&i.CurrencyID,
		&i.Price,
		&i.StatusID,
		&i.CreateUser,
		&i.ModifyUser,
		&i.CreatedAt,
		&i.ModifiedAt,
	)
	return i, err
}

const listProducts = `-- name: ListProducts :many
SELECT id, title, short_name, description, sms_noti, email_noti, call_noti, image, currency_id, price, status_id, create_user, modify_user, created_at, modified_at FROM cz_products
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListProductsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListProducts(ctx context.Context, arg ListProductsParams) ([]CzProduct, error) {
	rows, err := q.db.Query(ctx, listProducts, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []CzProduct{}
	for rows.Next() {
		var i CzProduct
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.ShortName,
			&i.Description,
			&i.SmsNoti,
			&i.EmailNoti,
			&i.CallNoti,
			&i.Image,
			&i.CurrencyID,
			&i.Price,
			&i.StatusID,
			&i.CreateUser,
			&i.ModifyUser,
			&i.CreatedAt,
			&i.ModifiedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateProducts = `-- name: UpdateProducts :one
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
RETURNING id, title, short_name, description, sms_noti, email_noti, call_noti, image, currency_id, price, status_id, create_user, modify_user, created_at, modified_at
`

type UpdateProductsParams struct {
	Title       string    `json:"title"`
	ShortName   string    `json:"short_name"`
	Description string    `json:"description"`
	SmsNoti     bool      `json:"sms_noti"`
	EmailNoti   bool      `json:"email_noti"`
	CallNoti    bool      `json:"call_noti"`
	Image       string    `json:"image"`
	CurrencyID  int32     `json:"currency_id"`
	Price       float64   `json:"price"`
	StatusID    int32     `json:"status_id"`
	ModifyUser  int64     `json:"modify_user"`
	ModifiedAt  time.Time `json:"modified_at"`
	ID          int32     `json:"id"`
}

func (q *Queries) UpdateProducts(ctx context.Context, arg UpdateProductsParams) (CzProduct, error) {
	row := q.db.QueryRow(ctx, updateProducts,
		arg.Title,
		arg.ShortName,
		arg.Description,
		arg.SmsNoti,
		arg.EmailNoti,
		arg.CallNoti,
		arg.Image,
		arg.CurrencyID,
		arg.Price,
		arg.StatusID,
		arg.ModifyUser,
		arg.ModifiedAt,
		arg.ID,
	)
	var i CzProduct
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.ShortName,
		&i.Description,
		&i.SmsNoti,
		&i.EmailNoti,
		&i.CallNoti,
		&i.Image,
		&i.CurrencyID,
		&i.Price,
		&i.StatusID,
		&i.CreateUser,
		&i.ModifyUser,
		&i.CreatedAt,
		&i.ModifiedAt,
	)
	return i, err
}
