package controller

import (
	"context"
	"github.com/Onelvay/HL-architecture/internal/dto"
	"github.com/Onelvay/HL-architecture/internal/entity"
	"github.com/Onelvay/HL-architecture/internal/repository"
	"github.com/google/uuid"
)

type OrderService struct {
	repo repository.OrderRepository
}

func NewOrderService(repo repository.OrderRepository) *OrderService {
	return &OrderService{
		repo: repo,
	}
}

func (o *OrderService) Create(ctx context.Context, req dto.OrderRequest) (res dto.OrderResponse) {
	id := uuid.New().String()
	order := entity.Order{
		ID:       id,
		UserId:   req.UserId,
		CourseId: req.CourseId,
	}

	err := o.repo.InsertOne(ctx, order)
	if err != nil {
		res.Status = 500
		res.Error = err
	}
	res.Status = 200
	return

}

func (o *OrderService) GetMany(ctx context.Context, userId string) (res []entity.Order, err error) {
	res, err = o.repo.GetManyById(context.Background(), userId)

	return
}

func (o *OrderService) AddReview(ctx context.Context, req dto.OrderReviewRequest) (err error) {
	review := entity.OrderReview{
		OrderID: req.OrderId,
		Comment: req.Comment,
		Rating:  req.Rating,
	}

	err = o.repo.AddReview(ctx, review)

	return
}
