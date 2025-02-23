package dto

import "github.com/CP-RektMart/pic-me-pls-backend/internal/model"

type TokenResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	Exp          int64  `json:"exp"`
}

type RegisterRequest struct {
	Provider string `json:"provider" validate:"required,provider"` // GOOGLE
	Role     string `json:"role" validate:"required,role"`         // CUSTOMER, PHOTOGRAPHER, ADMIN
	IDToken  string `json:"idToken" validate:"required"`
}

type RegisterResponse struct {
	TokenResponse
	User UserResponse `json:"user"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refreshToken" validate:"required"`
}

type LoginRequest struct {
	Provider model.Provider `json:"provider" required:"true" validate:"required,provider"`
	IDToken  string         `json:"idToken" required:"true" validate:"required"`
}

type LoginResponse struct {
	TokenResponse
	User UserResponse `json:"user"`
}
