package service

import (
	"context"
	"github.com/Onelvay/HL-architecture/internal/dto"
	"github.com/Onelvay/HL-architecture/internal/repository"
)

type authService struct {
	repo repository.AuthRepository
}

func NewAuthService(repo repository.AuthRepository) AuthService {
	return &authService{repo: repo}
}

func (a *authService) SignIn(ctx context.Context, req dto.SignInRequest) (res dto.SignInResponse, err error) {

	return
}

func (a *authService) SignUp(ctx context.Context, req dto.SignUpRequest) (res dto.SignUpResponse, err error) {
	err = a.repo.SignUp(ctx, req)
	return
}
