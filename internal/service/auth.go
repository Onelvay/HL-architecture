package service

import (
	"context"
	"github.com/Onelvay/HL-architecture/internal/dto"
	"github.com/Onelvay/HL-architecture/internal/entity"
	"github.com/Onelvay/HL-architecture/internal/repository"
	"github.com/Onelvay/HL-architecture/pkg/hasher"
	"github.com/dgrijalva/jwt-go/v4"
	"time"
)

type authService struct {
	repo repository.AuthRepository
}

const (
	expireDuration = 15 * time.Minute
)

func NewAuthService(repo repository.AuthRepository) AuthService {
	return &authService{repo: repo}
}

func (a *authService) SignIn(ctx context.Context, req dto.SignInRequest) (res dto.SignInResponse, err error) {
	user := entity.User{
		Email:    req.Email,
		Password: req.Password,
	}
	hasher.Password(&user)

	user, err = a.repo.SignIn(ctx, user)

	if err != nil {
		return
	}
	res.AccessToken, err = generateToken(user.Name)

	return
}

func (a *authService) SignUp(ctx context.Context, req dto.SignUpRequest) (res dto.SignUpResponse, err error) {
	user := entity.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	hasher.Password(&user)

	err = a.repo.SignUp(ctx, user)
	if err != nil {
		return res, err
	}

	res.AccessToken, err = generateToken(user.Name)
	return
}

func generateToken(username string) (accessToken string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &dto.Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(expireDuration)),
			IssuedAt:  jwt.At(time.Now()),
		},
		Username: username,
	})

	var key []byte
	accessToken, err = token.SignedString(key)
	if err != nil {
		return
	}
	return
}
