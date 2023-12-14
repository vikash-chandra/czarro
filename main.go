package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/czarro/api"
	db "github.com/czarro/db/sqlc"
	"github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
)

const (
	dbSource = "postgresql://root:secret@localhost:5432/carzorro?sslmode=disable"
	addresss = "127.0.0.1:8080"
)

func main() {
	conn, err := pgx.Connect(context.Background(), dbSource)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())
	queries := db.New(conn)
	server := api.NewServer(queries)
	err = server.Start(addresss)
	if err != nil {
		log.Fatal(err)
	}
}
