package service

import (
	"context"
	"github.com/Onelvay/HL-architecture/internal/dto"
)

type CourseService interface {
	GetRowById(ctx context.Context, id string) (res dto.CourseResponse, err error)
}

type AuthService interface {
	SignUp(ctx context.Context, req dto.SignUpRequest) (res dto.SignUpResponse, err error)
	SignIn(ctx context.Context, req dto.SignInRequest) (res dto.SignInResponse, err error)
}
