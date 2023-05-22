package postgres

import (
	"context"
	"github.com/Onelvay/HL-architecture/internal/entity"
	"github.com/jmoiron/sqlx"
)

type CourseRepository struct {
	db *sqlx.DB
}

func NewCourseRepository(db *sqlx.DB) *CourseRepository {
	return &CourseRepository{db: db}
}

func (p *CourseRepository) GetOne(ctx context.Context, id string) (data entity.Course, err error) {
	query := `SELECT id,name,price FROM products WHERE id=$1`
	args := []any{id}

	err = p.db.GetContext(ctx, &data, query, args...)
	return
}
