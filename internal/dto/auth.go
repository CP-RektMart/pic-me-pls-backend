package dto

type RegisterRequest struct {
	Provider string `json:"provider" validate:"required,provider"` // GOOGLE
	Role     string `json:"role" validate:"required,role"`         // CUSTOMER, PHOTOGRAPHER, ADMIN
	IDToken  string `json:"idToken" validate:"required"`
}

type TokenResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	Exp          int64  `json:"exp"`
}

type RegisterResponse struct {
	TokenResponse
	User UserResponse `json:"user"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refreshToken" validate:"required"`
}
