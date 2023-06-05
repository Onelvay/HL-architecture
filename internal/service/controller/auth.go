package controller

import (
	"context"
	"github.com/Onelvay/HL-architecture/internal/domain/auth"
	"github.com/Onelvay/HL-architecture/internal/domain/user"
	"github.com/Onelvay/HL-architecture/internal/repository"
	"github.com/Onelvay/HL-architecture/pkg/hasher"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/google/uuid"
	"time"
)

type AuthService struct {
	repo repository.AuthRepository
}

const (
	expireDuration = 18000 * time.Minute // временно
)

func NewAuthService(repo repository.AuthRepository) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (a *AuthService) SignIn(ctx context.Context, req auth.SignInRequest) (res auth.SignInResponse, err error) {
	user := user.User{
		Email:    req.Email,
		Password: req.Password,
	}
	hasher.Password(&user)

	user, err = a.repo.SignIn(ctx, user)

	if err != nil {
		return
	}
	res.AccessToken, err = generateToken(user.ID)

	return
}

func (a *AuthService) SignUp(ctx context.Context, req auth.SignUpRequest) (res auth.SignUpResponse, err error) {
	user := user.User{
		ID:       uuid.New().String(),
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	hasher.Password(&user)

	err = a.repo.SignUp(ctx, user)
	if err != nil {
		return res, err
	}

	res.AccessToken, err = generateToken(user.ID)
	return
}

func generateToken(userId string) (accessToken string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &auth.Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(expireDuration)),
			IssuedAt:  jwt.At(time.Now()),
		},
		UserId: userId,
	})

	var key []byte
	accessToken, err = token.SignedString(key)
	if err != nil {
		return
	}
	return
}
