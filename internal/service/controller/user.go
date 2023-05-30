package controller

import (
	"context"
	"github.com/Onelvay/HL-architecture/internal/dto"
	"github.com/Onelvay/HL-architecture/internal/repository"
	"github.com/Onelvay/HL-architecture/pkg/logger"
)

type UserService struct {
	userRepo   repository.UserRepository
	orderRepo  repository.UserOrderRepository
	courseRepo repository.CourseRepository
}

func NewUserService(user repository.UserRepository, order repository.UserOrderRepository, course repository.CourseRepository) *UserService {
	return &UserService{
		userRepo:   user,
		orderRepo:  order,
		courseRepo: course,
	}
}

func (u *UserService) GetCourses(ctx context.Context, userID string) (res []dto.CourseResponse, err error) {
	orders, err := u.orderRepo.GetManyById(ctx, userID)
	//временно
	for _, v := range orders {
		course, err := u.courseRepo.GetOne(ctx, v.CourseId)

		if err != nil {
			logger.Error(err)
		} else {
			res = append(res,
				dto.CourseResponse{
					ID:          course.ID,
					Name:        course.Name,
					Description: course.Description,
					ImgURL:      course.ImgURL,
				},
			)
		}
	}

	return
}
