package postgres

import (
	"context"
	"errors"
	"github.com/Onelvay/HL-architecture/internal/dto"
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

	return
}

func (o *OrderRepository) exist(ctx context.Context, userid, id string) error {
	query := `SELECT order_id,user_id,course_id from orders where user_id = $1 and course_id=$2 limit 1 `
	res, err := o.db.Exec(query, userid, id)

	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if rows != 0 {
		return errors.New("exist")
	}
	return nil
}

func (o *OrderRepository) AddReview(ctx context.Context, req entity.OrderReview) (err error) {
	query := "INSERT INTO reviews (order_id, comment,rating) VALUES (:id, :comment,:rating)"

	_, err = o.db.NamedExec(query, map[string]interface{}{
		"id":      req.OrderID,
		"comment": req.Comment,
		"rating":  req.Rating,
	})

	return
}

func (o *OrderRepository) GetAllReviews(ctx context.Context) (orders []dto.ReviewResponse, err error) {
	query := `select u.email as user_name,c.name as course_name, r.rating,r.comment,r.created_at from orders 
    inner join courses c on orders.course_id = c.id 
    inner join users u on u.id = orders.user_id 
    inner join reviews r on orders.order_id = r.order_id`

	err = o.db.Select(&orders, query)

	return
}
