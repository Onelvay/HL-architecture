package controller

import (
	"context"
	courseD "github.com/Onelvay/HL-architecture/internal/domain/course"
	"github.com/Onelvay/HL-architecture/internal/repository"
	"github.com/Onelvay/HL-architecture/pkg/loggerlocal"
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

func (u *UserService) GetCourses(ctx context.Context, userID string) (res []courseD.UserCourseResponse, err error) {
	orders, err := u.orderRepo.GetManyById(ctx, userID)
	//временно
	for _, order := range orders {
		course, err := u.courseRepo.GetOne(ctx, order.CourseId)

		if err != nil {
			loggerlocal.Error(err)
		} else {
			res = append(res,
				courseD.UserCourseResponse{
					ID:          course.ID,
					Name:        course.Name,
					Description: course.Description,
					ImgURL:      course.ImgURL,
					OrderId:     order.ID,
				},
			)
		}
	}

	return
}
