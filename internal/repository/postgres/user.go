package postgres

import (
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

//func (u *UserRepository) GetCourses(ctx context.Context, userID string) (courses []entity.Course, err error) {
//	return
//}
