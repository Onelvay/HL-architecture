package controller

import (
	"context"
	"github.com/Onelvay/HL-architecture/internal/dto"
	"github.com/Onelvay/HL-architecture/internal/entity"
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
	expireDuration = 90 * time.Minute // временно
)

func NewAuthService(repo repository.AuthRepository) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (a *AuthService) SignIn(ctx context.Context, req dto.SignInRequest) (res dto.SignInResponse, err error) {
	user := entity.User{
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

func (a *AuthService) SignUp(ctx context.Context, req dto.SignUpRequest) (res dto.SignUpResponse, err error) {
	user := entity.User{
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
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &dto.Claims{
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
