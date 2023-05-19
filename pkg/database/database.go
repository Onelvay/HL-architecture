package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

type Order struct {
	Id     string `db:"id"`
	UserId string `db:"user_id"`
}

func New() {
	db, err := sqlx.Connect("postgres", "host=localhost user=postgres dbname=bookstore password=Adg12332, sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

}
