package postgres

import (
	"context"
	"github.com/Onelvay/HL-architecture/internal/entity"
	"github.com/jmoiron/sqlx"
)

type ProductRepository struct {
	db *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (p *ProductRepository) GetOne(ctx context.Context, id string) (data entity.Product, err error) {
	query := `SELECT id,name,price FROM products WHERE id=$1`
	args := []any{id}

	err = p.db.GetContext(ctx, &data, query, args...)
	return
}
