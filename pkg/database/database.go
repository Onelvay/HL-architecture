package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func New() (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", "postgres://abay:s60VgoNFJN351yzxJueod9ctqiVHaQ0p@dpg-ci026nu4dad35tqp0ec0-a/aray")
	if err != nil {
		return db, err
	}

	if err := migrates(); err != nil {
		panic(err)
	}

	return db, nil
}
