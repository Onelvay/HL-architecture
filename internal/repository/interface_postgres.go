package repository

import (
	"context"
	"github.com/Onelvay/HL-architecture/internal/dto"
	"github.com/Onelvay/HL-architecture/internal/entity"
)

type ProductRepository interface {
	GetOne(ctx context.Context, id string) (data entity.Product, err error)
}
type AuthRepository interface {
	SignUp(ctx context.Context, req dto.SignUpRequest) (err error)
	SignIn(ctx context.Context, req dto.SignInRequest) (err error)
}
