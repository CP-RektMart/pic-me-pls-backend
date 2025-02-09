package auth

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/database"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/services/jwt"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	store          *database.Store
	validate       *validator.Validate
	jwtService     *jwt.Service
	googleClientID string
}

func NewHandler(store *database.Store, validate *validator.Validate, jwtService *jwt.Service, googleClientID string) *Handler {
	return &Handler{
		store:          store,
		validate:       validate,
		jwtService:     jwtService,
		googleClientID: googleClientID,
	}
}
