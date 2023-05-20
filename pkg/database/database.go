package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Order struct {
	Id     string `db:"id"`
	UserId string `db:"user_id"`
}

func New() (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", "host=localhost user=postgres dbname=products password=Adg12332, sslmode=disable")
	if err != nil {
		return db, err
	}
	return db, nil
}
