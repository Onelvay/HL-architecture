package service

import "github.com/Onelvay/HL-architecture/internal/repository/postgres"

type Dependencies struct {
	ProductRepo *postgres.ProductRepository
}

type Service struct {
	Product ProductService
}

func New(d Dependencies) Service {
	return Service{
		Product: d.ProductRepo,
	}
}
