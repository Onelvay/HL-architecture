package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func New() (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", "postgres://abay:8g2ct20M3j602Kgel5BL8BRbFRUwIbqC@dpg-ci02ko33cv20nhr4f44g-a/aray_29yf")
	if err != nil {
		return db, err
	}

	if err := migrates(); err != nil {
		panic(err)
	}

	return db, nil
}
