package dto

type LoginRequest struct {
	Provider string `json:"provider" validate:"required,provider"`
	Role     string `json:"role" validate:"required,role"`
	IDToken  string `json:"idToken" validate:"required"`
}

type TokenResponse struct {
	AcessToken   string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	Exp          int    `json:"exp"`
}

type LoginResponse struct {
	TokenResponse
	User BaseUserDTO `json:"user"`
}
