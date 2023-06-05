package postgres

import (
	"context"
	"github.com/Onelvay/HL-architecture/internal/domain/course"
	"github.com/jmoiron/sqlx"
)

type CourseRepository struct {
	db *sqlx.DB
}

func NewCourseRepository(db *sqlx.DB) *CourseRepository {
	return &CourseRepository{db: db}
}

func (p *CourseRepository) GetOne(ctx context.Context, id string) (data course.Course, err error) {
	query := `SELECT id,name,description,imgurl FROM courses WHERE id=$1`
	args := []any{id}

	err = p.db.GetContext(ctx, &data, query, args...)
	return
}

func (p *CourseRepository) GetMany(ctx context.Context) (data []course.Course, err error) {
	query := `SELECT id,name,description,imgurl FROM courses`

	err = p.db.SelectContext(ctx, &data, query)
	return
}
