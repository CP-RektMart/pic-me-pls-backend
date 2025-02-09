package auth

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/database"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/services/jwt"
	"github.com/go-playground/validator/v10"
)

type Config struct {
	GoogleClientID string `env:"GOOGLE_CLIENT_ID"`
}

type Handler struct {
	config     Config
	store      *database.Store
	validate   *validator.Validate
	jwtService *jwt.Service
}

func NewHandler(config Config, store *database.Store, validate *validator.Validate, jwtService *jwt.Service) *Handler {
	return &Handler{
		config:     config,
		store:      store,
		validate:   validate,
		jwtService: jwtService,
	}
}
