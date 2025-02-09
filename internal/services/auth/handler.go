package auth

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/database"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/jwt"
	"github.com/go-playground/validator/v10"
)

type Config struct {
	GoogleClientID string `env:"GOOGLE_CLIENT_ID"`
}

type Handler struct {
	store        *database.Store
	validate  *validator.Validate
	JWTConfig jwt.Config
	config    Config
}

func NewHandler(store *database.Store, validate *validator.Validate, jwtConfig jwt.Config, config Config) *Handler {
	return &Handler{
		store:        store,
		validate:  validate,
		JWTConfig: jwtConfig,
		config:    config,
	}
}
