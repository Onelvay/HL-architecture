package controller

import (
	"context"
	"github.com/Onelvay/HL-architecture/internal/dto"
	"github.com/Onelvay/HL-architecture/internal/entity"
	"github.com/Onelvay/HL-architecture/internal/repository"
	"github.com/Onelvay/HL-architecture/pkg/logger"
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

func (o *OrderService) CreateRow(ctx context.Context, req dto.OrderRequest) (res dto.OrderResponse) {
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

func (o *OrderService) GetAllReviews(ctx context.Context) (orders []dto.ReviewResponse, err error) {
	orders, err = o.repo.GetAllReviews(ctx)
	if err != nil {
		logger.Error(err)
	}
	return
}

func (o *OrderService) DeleteByOrderId(ctx context.Context, req dto.OrderDeleteRequest) (err error) {
	orderId := req.OrderId
	err = o.repo.DeleteRow(ctx, orderId)
	return
}
