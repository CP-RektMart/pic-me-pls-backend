package chat

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/database"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/middlewares/authentication"
)

type Handler struct {
	store          *database.Store
	chatSystem     *ChatSystem
	authentication authentication.AuthMiddleware
}

func NewHandler(store *database.Store, authentication authentication.AuthMiddleware) *Handler {
	return &Handler{
		store:          store,
		chatSystem:     NewChatSystem(),
		authentication: authentication,
	}
}
