package postgres

import (
	"context"
	"github.com/Onelvay/HL-architecture/internal/domain/user"
	"github.com/jmoiron/sqlx"
)

type AuthRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (a *AuthRepository) SignIn(ctx context.Context, req user.User) (user user.User, err error) {
	query := `
		SELECT id,name,email,password FROM users WHERE email=$1 and password=$2`

	args := []any{req.Email, req.Password}

	err = a.db.GetContext(ctx, &user, query, args...)

	return
}

func (a *AuthRepository) SignUp(ctx context.Context, req user.User) (err error) {
	query := "INSERT INTO users (id,name, email,password) VALUES (:id, :name, :email,:password)"

	_, err = a.db.NamedExec(query, map[string]interface{}{
		"id":       req.ID,
		"name":     req.Name,
		"email":    req.Email,
		"password": req.Password,
	})
	return
}
