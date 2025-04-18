package admin

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/database"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/middlewares/authentication"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	store          *database.Store
	authMiddleware authentication.AuthMiddleware
	validate       *validator.Validate
}

func NewHandler(
	store *database.Store,
	authMiddleware authentication.AuthMiddleware,
	validate *validator.Validate,
) *Handler {
	return &Handler{
		store:          store,
		authMiddleware: authMiddleware,
		validate:       validate,
	}
}
