package service

import (
	"context"
	"github.com/Onelvay/HL-architecture/internal/dto"
	"github.com/Onelvay/HL-architecture/internal/repository"
)

type courseService struct {
	repo repository.CourseRepository
}

func NewCourseService(repo repository.CourseRepository) CourseService {
	return &courseService{repo}
}

func (p *courseService) GetRowById(ctx context.Context, id string) (res dto.CourseResponse, err error) {
	product, err := p.repo.GetOne(ctx, id)

	res = dto.ParseFromCourse(product)
	return
}
