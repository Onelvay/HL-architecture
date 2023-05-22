package service

import "github.com/Onelvay/HL-architecture/internal/repository"

type Dependencies struct {
	CourseRepo repository.CourseRepository
	AuthRepo   repository.AuthRepository
}

type Service struct {
	Course CourseService
	Auth   AuthService
}

func New(d Dependencies) Service {
	return Service{
		Course: NewCourseService(d.CourseRepo),
		Auth:   NewAuthService(d.AuthRepo),
	}
}
