package service

import (
	"github.com/Onelvay/HL-architecture/internal/repository"
	"github.com/Onelvay/HL-architecture/internal/service/controller"
)

type Dependencies struct {
	CourseRepo repository.CourseRepository
	AuthRepo   repository.AuthRepository
	OrderRepo  repository.OrderRepository
}

type Service struct {
	Course CourseService
	Auth   AuthService
	Order  OrderService
}

func New(d Dependencies) Service {
	return Service{
		Course: controller.NewCourseService(d.CourseRepo),
		Auth:   controller.NewAuthService(d.AuthRepo),
		Order:  controller.NewOrderService(d.OrderRepo),
	}
}
