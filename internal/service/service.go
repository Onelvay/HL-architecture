package service

import "github.com/Onelvay/HL-architecture/internal/repository"

type Dependencies struct {
	ProductRepo repository.ProductRepository
}

type Service struct {
	Product ProductService
}

func New(d Dependencies) Service {
	return Service{
		Product: NewProductService(d.ProductRepo),
	}
}
