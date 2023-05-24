package service

import "github.com/Onelvay/HL-architecture/internal/repository"

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
		Course: newCourseService(d.CourseRepo),
		Auth:   newAuthService(d.AuthRepo),
		Order:  newOrderService(d.OrderRepo),
	}
}
