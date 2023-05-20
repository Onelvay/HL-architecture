package service

import (
	"context"
	"github.com/Onelvay/HL-architecture/internal/dto"
	"github.com/Onelvay/HL-architecture/internal/repository"
)

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{repo}
}

func (p *productService) GetRowById(ctx context.Context, id string) (res dto.ProductResponse, err error) {
	product, err := p.repo.GetOne(ctx, id)

	res = dto.ParseFromProduct(product)
	return
}
