package model

import "github.com/golang-jwt/jwt/v5"

type CustomClaim struct {
	Role   string
	UserID uint
	jwt.RegisteredClaims
}
