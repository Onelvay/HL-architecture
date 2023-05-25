package service

import (
	"github.com/Onelvay/HL-architecture/internal/repository"
	"github.com/Onelvay/HL-architecture/internal/service/controller"
)

type Dependencies struct {
	CourseRepo repository.CourseRepository
	AuthRepo   repository.AuthRepository
	OrderRepo  repository.OrderRepository
	UserRepo   repository.UserRepository
}

type Service struct {
	Course CourseService
	Auth   AuthService
	Order  OrderService
	User   UserService
}

func New(d Dependencies) Service {
	return Service{
		Course: controller.NewCourseService(d.CourseRepo),
		Auth:   controller.NewAuthService(d.AuthRepo),
		Order:  controller.NewOrderService(d.OrderRepo),
		User:   controller.NewUserService(d.UserRepo),
	}
}
