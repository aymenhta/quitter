package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func SetupDb(dbUrl string) (*pgxpool.Pool, error) {
	dbpool, err := pgxpool.New(context.Background(), dbUrl)
	if err != nil {
		return nil, err
	}
	if err = dbpool.Ping(context.Background()); err != nil {
		return nil, err
	}
	return dbpool, nil
}

func TestDb(db *pgxpool.Pool) {
	var feedback string
	err := db.QueryRow(context.Background(), "select 'CONNECTED TO DB SUCCSSEFULLY!'").Scan(&feedback)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	log.Println(feedback)
}
