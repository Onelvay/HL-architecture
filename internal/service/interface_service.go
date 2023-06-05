package service

import (
	"context"
	"github.com/Onelvay/HL-architecture/internal/domain/auth"
	"github.com/Onelvay/HL-architecture/internal/domain/course"
	"github.com/Onelvay/HL-architecture/internal/domain/order"
)

type CourseService interface {
	GetRowById(ctx context.Context, id string) (res course.CourseResponse, err error)
	GetMany(ctx context.Context) (res []course.CourseResponse, err error)
}

type AuthService interface {
	SignUp(ctx context.Context, req auth.SignUpRequest) (res auth.SignUpResponse, err error)
	SignIn(ctx context.Context, req auth.SignInRequest) (res auth.SignInResponse, err error)
}

type OrderService interface {
	CreateRow(ctx context.Context, req order.OrderRequest) (res order.OrderResponse)
	GetMany(ctx context.Context, userId string) (res []order.Order, err error)
	AddReview(ctx context.Context, req order.OrderReviewRequest) (err error)
	GetAllReviews(ctx context.Context) (orders []order.ReviewResponse, err error)
	DeleteByOrderId(ctx context.Context, req order.OrderDeleteRequest) error
}

type UserService interface {
	GetCourses(ctx context.Context, userID string) ([]course.UserCourseResponse, error)
}
