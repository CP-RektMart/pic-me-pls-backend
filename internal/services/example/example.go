package example

import "github.com/CP-RektMart/pic-me-pls-backend/internal/database"

type Handler struct {
	db *database.Store
}

func NewHandler(db *database.Store) *Handler {
	return &Handler{db: db}
}
