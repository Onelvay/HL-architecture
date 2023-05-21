package service

import (
	"context"
	"github.com/Onelvay/HL-architecture/internal/dto"
	"github.com/Onelvay/HL-architecture/internal/entity"
	"github.com/Onelvay/HL-architecture/internal/repository"
)

type authService struct {
	repo repository.AuthRepository
}

func NewAuthService(repo repository.AuthRepository) AuthService {
	return &authService{repo: repo}
}

func (a *authService) SignIn(ctx context.Context, req dto.SignInRequest) (res dto.SignInResponse, err error) {
	user := entity.User{Email: req.Email, Password: req.Password}
	_, err = a.repo.SignIn(ctx, user)
	if err != nil {
		return
	}
	res.AccessToken = "ASDASDADS"
	return
}

func (a *authService) SignUp(ctx context.Context, req dto.SignUpRequest) (res dto.SignUpResponse, err error) {
	data := convertReq(req)
	err = a.repo.SignUp(ctx, data)
	if err != nil {
		return res, err
	}
	res.AccessToken = "asdasdasd"
	return
}

func convertReq(req dto.SignUpRequest) entity.User {
	return entity.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
}
