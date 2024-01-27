package main

import (
	"context"
	"log"

	"github.com/czarro/api"
	db "github.com/czarro/db/sqlc"
	"github.com/czarro/logger"
	"github.com/czarro/util"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}
	logger.CZLoggerInit(config)
	connPool, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal(err)
	}
	queries := db.NewStore(connPool)
	server := api.NewServer(queries, config)
	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal(err)
	}
}
