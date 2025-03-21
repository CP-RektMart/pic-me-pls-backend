package message

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/chat"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/database"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/middlewares/authentication"
)

type Handler struct {
	store          *database.Store
	chatSystem     *chat.Server
	authentication authentication.AuthMiddleware
}

func NewHandler(store *database.Store, authentication authentication.AuthMiddleware, chatSystem *chat.Server) *Handler {
	return &Handler{
		store:          store,
		chatSystem:     chatSystem,
		authentication: authentication,
	}
}
