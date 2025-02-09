package server

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/middlewares/authentication"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/services/auth"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/services/example"
)

func (s *Server) RegisterRoutes(
	authMiddleware authentication.AuthMiddleware,
	exampleHandler *example.Handler,
	authHandler *auth.Handler,
) {
	api := s.app.Group("/api")
	v1 := api.Group("/v1")

	// auth
	auth := v1.Group("/auth")
	auth.Post("/login", authHandler.HandleLogin)

	// example
	example := v1.Group("/example")
	example.Post("/upload", exampleHandler.HandlerUploadExample)
}
