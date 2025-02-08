package dto

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
)

type LoginRequest struct {
	Provider string `json:"provider" validate:"required,oneof GOOGLE FACEBOOK APPLE"`
	Role     string `json:"role" validate:"required,oneof CUSTOMER PHOTOGRAPHER ADMIN"`
	IDToken  string `json:"idToken" validate:"required"`
}

type TokenResponse struct {
	AcessToken   string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	Exp          int    `json:"exp"`
}

type LoginResponse struct {
	TokenResponse
	User model.User `json:"user"`
}
