package repository

import (
	"context"
	"github.com/Onelvay/HL-architecture/internal/domain/course"
	"github.com/Onelvay/HL-architecture/internal/domain/order"
	"github.com/Onelvay/HL-architecture/internal/domain/user"
)

type CourseRepository interface {
	GetOne(ctx context.Context, id string) (data course.Course, err error)
	GetMany(ctx context.Context) (data []course.Course, err error)
}

type AuthRepository interface {
	SignIn(ctx context.Context, req user.User) (user user.User, err error)
	SignUp(ctx context.Context, req user.User) (err error)
}

type OrderRepository interface {
	InsertOne(ctx context.Context, req order.Order) (err error)
	GetManyById(ctx context.Context, userId string) (orders []order.Order, err error)
	AddReview(ctx context.Context, req order.OrderReview) (err error)
	GetAllReviews(ctx context.Context) (orders []order.ReviewResponse, err error)
	DeleteRow(ctx context.Context, orderId string) (err error)
}

type UserOrderRepository interface {
	GetManyById(ctx context.Context, userId string) (orders []order.Order, err error)
	InsertOne(ctx context.Context, req order.Order) (err error)
	AddReview(ctx context.Context, req order.OrderReview) (err error)
	DeleteRow(ctx context.Context, orderId string) (err error)
}

type UserRepository interface {
}
