package db

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	"log"
	"os"
	"testing"
	"tutorial.sqlc.dev/app/util"
)

var testQueries *Queries
var connPool *pgxpool.Pool

func testMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	connPool, err = pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	testQueries = New(connPool)

	os.Exit(m.Run())
}
