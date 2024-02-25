package main

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"tutorial.sqlc.dev/app/util"

	_ "github.com/lib/pq"
	"tutorial.sqlc.dev/app/api"
	db "tutorial.sqlc.dev/app/db/pgx"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	connPool, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(connPool)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
