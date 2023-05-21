package database

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func migrates() error {
	url := "postgres://postgres:Adg12332,@localhost:5432/aray?sslmode=disable"
	m, err := migrate.New("file://migrations", url)
	if err != nil {
		return err
	}

	if err = m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}
	//if err = m.Drop(); err != nil && err != migrate.ErrNoChange {
	//	return err
	//}

	return nil
}
