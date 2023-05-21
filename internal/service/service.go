package service

import "github.com/Onelvay/HL-architecture/internal/repository"

type Dependencies struct {
	ProductRepo repository.ProductRepository
	AuthRepo    repository.AuthRepository
}

type Service struct {
	Product ProductService
	Auth    AuthService
}

func New(d Dependencies) Service {
	return Service{
		Product: NewProductService(d.ProductRepo),
		Auth:    NewAuthService(d.AuthRepo),
	}
}
