package jwt

import (
	"fmt"
	"time"

	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/cockroachdb/errors"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(userID uint, role string, secret string, duration int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, model.CustomClaim{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(duration) * time.Second)),
		},
	})

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", errors.Wrap(err, "Failed to sign JWT")
	}
	return tokenString, nil
}

func ValidateJWT(token string, secret string) (*model.CustomClaim, error) {
	t, err := jwt.ParseWithClaims(token, &model.CustomClaim{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, fmt.Errorf("Failed to parse JWT: %w", err)
	}

	return t.Claims.(*model.CustomClaim), nil
}
