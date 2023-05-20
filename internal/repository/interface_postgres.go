package repository

import (
	"context"
	"github.com/Onelvay/HL-architecture/internal/entity"
)

type ProductRepository interface {
	GetOne(ctx context.Context, id string) (data entity.Product, err error)
}
