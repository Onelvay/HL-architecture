package dto

import "github.com/dgrijalva/jwt-go/v4"

type Claims struct {
	jwt.StandardClaims
	UserId string `json:"user_idr"`
}
