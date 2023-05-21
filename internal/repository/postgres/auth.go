package postgres

import (
	"context"
	"github.com/Onelvay/HL-architecture/internal/dto"
	"github.com/Onelvay/HL-architecture/internal/entity"
	"github.com/jmoiron/sqlx"
)

type AuthRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (a *AuthRepository) SignIn(ctx context.Context, req dto.SignInRequest) (user entity.User, err error) {
	query := `
		SELECT email,password FROM users WHERE email=$1 and password=$2`

	args := []any{req.Email, req.Password}

	err = a.db.GetContext(ctx, &user, query, args...)

	return
	return
}
func (a *AuthRepository) SignUp(ctx context.Context, req dto.SignUpRequest) (err error) {
	query := "INSERT INTO users (name, email,password) VALUES (:name, :email,:password)"

	_, err = a.db.NamedExec(query, map[string]interface{}{
		"name":     req.Name,
		"email":    req.Email,
		"password": req.Password,
	})
	return
}
