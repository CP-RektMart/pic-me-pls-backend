package admin

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/database"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/middlewares/authentication"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	store          *database.Store
	validate       *validator.Validate
	authMiddleware authentication.AuthMiddleware
}

func NewHandler(store *database.Store, validate *validator.Validate, authMiddleware authentication.AuthMiddleware) *Handler {
	return &Handler{
		store:          store,
		validate:       validate,
		authMiddleware: authMiddleware,
	}
}
