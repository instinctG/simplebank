package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	db "tutorial.sqlc.dev/app/db/pgx"
)

const dbURL = "postgresql://postgres:postgres@localhost:5432/simple_bank?sslmode=disable"

var q *db.Queries

func main() {

	connPool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	q = db.New(connPool)
	arg := db.CreateAccountParams{
		Owner:    "Tom",
		Balance:  100,
		Currency: "USD",
	}
	account, err := q.CreateAccount(context.Background(), arg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(account)
}
