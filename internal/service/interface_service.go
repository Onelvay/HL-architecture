package service

import (
	"context"
	"github.com/Onelvay/HL-architecture/internal/dto"
)

type ProductService interface {
	GetRowById(ctx context.Context, id string) (res dto.ProductResponse, err error)
}
