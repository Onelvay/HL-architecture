package service

import "github.com/Onelvay/HL-architecture/internal/repository/postgres"

type ProductService interface {
}

type productService struct {
	repo postgres.ProductRepository
}

func NewProductService(repo postgres.ProductRepository) *productService {
	return &productService{repo}
}
