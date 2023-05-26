package postgres

import (
	"context"
	"errors"
	"fmt"
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

func (o *OrderRepository) InsertOne(ctx context.Context, req entity.Order) (err error) {
	query := "INSERT INTO orders (order_id, user_id,course_id) VALUES (:id, :user_id,:course_id)"

	if err = o.exist(ctx, req.UserId, req.CourseId); err != nil {
		return err
	}
	_, err = o.db.NamedExec(query, map[string]interface{}{
		"id":        req.ID,
		"user_id":   req.UserId,
		"course_id": req.CourseId,
	})

	return
}

func (o *OrderRepository) GetManyById(ctx context.Context, userId string) (orders []entity.Order, err error) {
	query := `SELECT order_id,user_id,course_id from orders where user_id = $1`

	err = o.db.Select(&orders, query, userId)
	fmt.Println(orders, userId)
	return
}

func (o *OrderRepository) exist(ctx context.Context, userid, id string) error {
	var order entity.Order

	query := `SELECT order_id,user_id,course_id from orders where user_id = $1 and course_id=$2`
	err := o.db.Select(&order, query, userid, id)

	if err != nil {
		if order.ID != "" {
			return errors.New("exist")
		} else {
			return err
		}
	}

	return nil
}
