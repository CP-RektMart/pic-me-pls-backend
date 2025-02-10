package auth

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/database"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/jwt"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/middlewares/authentication"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	store          *database.Store
	validate       *validator.Validate
	jwtService     *jwt.JWT
	authMiddleware authentication.AuthMiddleware
	googleClientID string
}

func NewHandler(store *database.Store, validate *validator.Validate, jwtService *jwt.JWT, authMiddleware authentication.AuthMiddleware, googleClientID string) *Handler {
	return &Handler{
		store:          store,
		validate:       validate,
		jwtService:     jwtService,
		authMiddleware: authMiddleware,
		googleClientID: googleClientID,
	}
}
