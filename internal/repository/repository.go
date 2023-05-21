package repository

import (
	"github.com/Onelvay/HL-architecture/internal/repository/postgres"
	"github.com/Onelvay/HL-architecture/pkg/database"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	Product ProductRepository
	Auth    AuthRepository
}

type Configuration func(r *Repository) (db *sqlx.DB, err error)

func PostgresRepository() Configuration {
	return func(r *Repository) (db *sqlx.DB, err error) {
		db, err = database.New()
		if err != nil {
			return
		}

		r.Product = postgres.NewProductRepository(db)
		r.Auth = postgres.NewAuthRepository(db)
		return
	}
}

func New(cfg Configuration) (*Repository, *sqlx.DB, error) {
	repo := &Repository{}
	db, err := cfg(repo)

	if err != nil {
		return nil, nil, err
	}
	return repo, db, nil
}
