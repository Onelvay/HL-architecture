package controller

import (
	"context"
	"github.com/Onelvay/HL-architecture/internal/domain/course"
	"github.com/Onelvay/HL-architecture/internal/repository"
)

type CourseService struct {
	repo repository.CourseRepository
}

func NewCourseService(repo repository.CourseRepository) *CourseService {
	return &CourseService{repo}
}

func (p *CourseService) GetRowById(ctx context.Context, id string) (res course.CourseResponse, err error) {
	courses, err := p.repo.GetOne(ctx, id)

	res = course.ParseFromCourse(courses)
	return
}
func (p *CourseService) GetMany(ctx context.Context) (res []course.CourseResponse, err error) {
	courses, err := p.repo.GetMany(ctx)
	if err != nil {
		return nil, err
	}

	for _, v := range courses {
		res = append(res, course.ParseFromCourse(v))
	}
	return
}
