package db

import (
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

// Store ...
type Store interface {
	Querier
}

// SQLStore provides all functions to execute db queries and transaction
type SQLStore struct {
	connPool *pgxpool.Pool
	*Queries
}

// NewStore creates a new store
func NewStore(connPool *pgxpool.Pool) SQLStore {
	return SQLStore{
		connPool: connPool,
		Queries:  New(connPool),
	}
}
