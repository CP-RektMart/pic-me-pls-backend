package chat

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/database"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/middlewares/authentication"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	store          *database.Store
	chatSystem     *ChatSystem
	authentication authentication.AuthMiddleware
	validate       *validator.Validate
}

func NewHandler(store *database.Store, authentication authentication.AuthMiddleware, chatSystem *ChatSystem, validate *validator.Validate) *Handler {
	return &Handler{
		store:          store,
		chatSystem:     chatSystem,
		authentication: authentication,
		validate:       validate,
	}
}
