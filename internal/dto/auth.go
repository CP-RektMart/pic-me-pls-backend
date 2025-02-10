package dto

type LoginRequest struct {
	Provider string `json:"provider" validate:"required,provider"` // GOOGLE
	Role     string `json:"role" validate:"required,role"`         // CUSTOMER, PHOTOGRAPHER, ADMIN
	IDToken  string `json:"idToken" validate:"required"`
}

type TokenResponse struct {
	AcessToken   string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	Exp          int64  `json:"exp"`
}

type LoginResponse struct {
	TokenResponse
	User BaseUserDTO `json:"user"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refreshToken" validate:"required"`
}
