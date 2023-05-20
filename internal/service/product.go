package service

import (
	"github.com/Onelvay/HL-architecture/internal/repository"
)

type productService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) ProductService {
	return &productService{repo}
}
