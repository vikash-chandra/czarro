// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package db

import (
	"context"
)

type Querier interface {
	CreateCustomer(ctx context.Context, arg CreateCustomerParams) (Customer, error)
	DeleteCustomer(ctx context.Context, id int64) error
	GetCustomer(ctx context.Context, id int64) (Customer, error)
	ListCustomers(ctx context.Context, arg ListCustomersParams) ([]Customer, error)
	UpdateCustomer(ctx context.Context, arg UpdateCustomerParams) (Customer, error)
}

var _ Querier = (*Queries)(nil)
