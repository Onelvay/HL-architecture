package repository

import (
	"context"
	"github.com/Onelvay/HL-architecture/internal/entity"
)

type CourseRepository interface {
	GetOne(ctx context.Context, id string) (data entity.Course, err error)
}

type AuthRepository interface {
	SignIn(ctx context.Context, req entity.User) (user entity.User, err error)
	SignUp(ctx context.Context, req entity.User) (err error)
}
