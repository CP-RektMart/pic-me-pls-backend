package example

import "github.com/CP-RektMart/pic-me-pls-backend/internal/database"

type Handler struct {
	store *database.Store
}

func NewHandler(store *database.Store) *Handler {
	return &Handler{store: store}
}
