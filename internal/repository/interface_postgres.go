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

type OrderRepository interface {
	InsertOne(ctx context.Context, req entity.Order) (err error)
	GetManyById(ctx context.Context, userId string) (orders []entity.Order, err error)
}
