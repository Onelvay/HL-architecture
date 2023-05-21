package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func New() (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", "host=localhost user=postgres dbname=aray password=Adg12332, sslmode=disable")
	if err != nil {
		return db, err
	}

	if err := migrates(); err != nil {
		panic(err)
	}

	return db, nil
}
