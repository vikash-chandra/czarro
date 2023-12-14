package db

import (
	_ "github.com/lib/pq"
)

// Store ...
type Store interface {
	Queries
}

// SQLStore provides all functions to execute db queries and transaction
type SQLStore struct {
	// db *pgx.Conn
	*Queries
}

// NewStore creates a new store
func NewStore(q *Queries) SQLStore {
	return SQLStore{
		Queries: q,
	}
}
