package service

import "github.com/Onelvay/HL-architecture/internal/repository/postgres"

type Dependencies struct {
	productRepo postgres.ProductRepository
}

type Service struct {
	Author ProductService
}

func New(d Dependencies) Service {
	return Service{
		d.productRepo,
	}
}
