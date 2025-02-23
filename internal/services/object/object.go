package object

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/database"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/storage"
)

type Handler struct {
	store  *database.Store
	config storage.Config
}

func NewHandler(store *database.Store, config storage.Config) *Handler {
	return &Handler{
		store:  store,
		config: config,
	}
}
