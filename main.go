package main

import (
	"context"
	"fmt"
	"log"

	"github.com/czarro/api"
	db "github.com/czarro/db/sqlc"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

const (
	dbSource = "postgresql://root:secret@localhost:5432/carzorro?sslmode=disable"
	addresss = "127.0.0.1:8080"
)

func main() {
	connPool, err := pgxpool.New(context.Background(), dbSource)
	if err != nil {
		fmt.Println(err.Error())
	}
	queries := db.NewStore(connPool)
	server := api.NewServer(queries)
	err = server.Start(addresss)
	if err != nil {
		log.Fatal(err)
	}
}