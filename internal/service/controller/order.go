package controller

import (
	"context"
	"github.com/Onelvay/HL-architecture/internal/domain/order"
	"github.com/Onelvay/HL-architecture/internal/repository"
	"github.com/Onelvay/HL-architecture/pkg/loggerlocal"
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

func (o *OrderService) CreateRow(ctx context.Context, req order.OrderRequest) (res order.OrderResponse) {
	id := uuid.New().String()
	order := order.Order{
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

func (o *OrderService) GetMany(ctx context.Context, userId string) (res []order.Order, err error) {
	res, err = o.repo.GetManyById(context.Background(), userId)

	return
}

func (o *OrderService) AddReview(ctx context.Context, req order.OrderReviewRequest) (err error) {
	review := order.OrderReview{
		OrderID: req.OrderId,
		Comment: req.Comment,
		Rating:  req.Rating,
	}

	err = o.repo.AddReview(ctx, review)

	return
}

func (o *OrderService) GetAllReviews(ctx context.Context) (orders []order.ReviewResponse, err error) {
	orders, err = o.repo.GetAllReviews(ctx)
	if err != nil {
		loggerlocal.Error(err)
	}
	return
}

func (o *OrderService) DeleteByOrderId(ctx context.Context, req order.OrderDeleteRequest) (err error) {
	orderId := req.OrderId
	err = o.repo.DeleteRow(ctx, orderId)
	return
}
