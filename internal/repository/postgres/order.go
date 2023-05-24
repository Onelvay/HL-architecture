package postgres

import (
	"context"
	"github.com/Onelvay/HL-architecture/internal/entity"
	"github.com/jmoiron/sqlx"
)

type OrderRepository struct {
	db *sqlx.DB
}

func NewOrderRepository(db *sqlx.DB) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}
func (o *OrderRepository) POST(ctx context.Context, req entity.Order) (err error) {
	query := "INSERT INTO orders (order_id, user_id,course_id) VALUES (:id, :user_id,:course_id)"

	_, err = o.db.NamedExec(query, map[string]interface{}{
		"id":        req.ID,
		"user_id":   req.UserId,
		"course_id": req.CourseId,
	})

	return
}
func (o *OrderRepository) GET(ctx context.Context, userId string) (orders []entity.Order, err error) {

	return
}
