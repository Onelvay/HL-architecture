package service

import (
	"context"
	"github.com/Onelvay/HL-architecture/internal/dto"
	"github.com/Onelvay/HL-architecture/internal/entity"
	"github.com/Onelvay/HL-architecture/internal/repository"
	"github.com/google/uuid"
)

type orderService struct {
	repo repository.OrderRepository
}

func newOrderService(repo repository.OrderRepository) OrderService {
	return &orderService{
		repo: repo,
	}
}

func (o *orderService) Create(ctx context.Context, req dto.OrderRequest) (res dto.OrderResponse) {
	id := uuid.New().String()
	order := entity.Order{
		ID:       id,
		UserId:   req.UserId,
		CourseId: req.CourseId,
	}

	err := o.repo.POST(ctx, order)
	if err != nil {
		res.Status = 500
		res.Error = err
	}
	res.Status = 200
	return

}
