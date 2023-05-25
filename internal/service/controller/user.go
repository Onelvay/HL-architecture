package controller

import (
	"context"
	"github.com/Onelvay/HL-architecture/internal/entity"
	"github.com/Onelvay/HL-architecture/internal/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (u *UserService) GetCourses(ctx context.Context) (courses []entity.Course, err error) {
	return
}
