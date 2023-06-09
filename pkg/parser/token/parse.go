package token

import (
	"errors"
	"fmt"
	"github.com/Onelvay/HL-architecture/internal/domain/auth"
	"github.com/dgrijalva/jwt-go/v4"
)

func Parse(accessToken string, signingKey []byte) (string, error) {
	token, err := jwt.ParseWithClaims(accessToken, &auth.Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return signingKey, nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(*auth.Claims); ok && token.Valid {
		return claims.UserId, nil
	}

	return "", errors.New("invalid auth token")
}
