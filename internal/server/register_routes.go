package server

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/middlewares/authentication"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/services/auth"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/services/category"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/services/example"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/services/gallery"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/services/message"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/services/photographer"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/services/review"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/services/user"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/services/verifycard"
)

func (s *Server) RegisterRoutes(
	authMiddleware authentication.AuthMiddleware,
	exampleHandler *example.Handler,
	authHandler *auth.Handler,
	userHandler *user.Handler,
	photographerHandler *photographer.Handler,
	galleryHandler *gallery.Handler,
	reviewHandler *review.Handler,
	categoryHandler *category.Handler,
	messageHandler *message.Handler,
	verifyCardHandler *verifycard.Handler,
) {
	api := s.app.Group("/api")
	v1 := api.Group("/v1")

	// example
	example := v1.Group("/example")
	example.Post("/upload", exampleHandler.HandlerUploadExample)

	// auth
	auth := v1.Group("/auth")
	auth.Post("/login", authHandler.HandleLogin)

	// profile
	// TODO: v1.Get("/me", authMiddleware.Auth, userHandler.HandlerUpdateProfile) add Middleware
	v1.Post("/me", userHandler.HandlerUpdateProfile) // add Middleware

	// verify citizen card
	photographer := v1.Group("/photographer")
	photographer.Post("/verify", verifyCardHandler.HandlerVerifyCard)      // add Middleware
	photographer.Patch("/reverify", verifyCardHandler.HandlerReVerifyCard) // add Middleware
}
