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
) {
	api := s.app.Group("/api")
	v1 := api.Group("/v1")

	// example
	example := v1.Group("/example")
	example.Post("/upload", exampleHandler.HandlerUploadExample)

	// auth
	auth := v1.Group("/auth")
	auth.Post("/login", authHandler.HandleLogin)
	auth.Post("/register", authHandler.HandleRegister)
	auth.Post("/refresh-token", authHandler.HandleRefreshToken)

	// user
	v1.Get("/me", authMiddleware.Auth, userHandler.HandleGetMe)
	v1.Patch("/me", authMiddleware.Auth, userHandler.HandleUpdateMe)

	// verify citizen card
	photographer := v1.Group("/photographer")
	photographer.Get("/citizen-card", authMiddleware.AuthPhotographer, photographerHandler.HandleGetCitizenCard)
	photographer.Post("/verify", authMiddleware.AuthPhotographer, photographerHandler.HandleVerifyCard)
	photographer.Patch("/reverify", authMiddleware.AuthPhotographer, photographerHandler.HandleReVerifyCard)

	auth.Post("/logout", authMiddleware.Auth, authHandler.HandleLogout)

	// gallery
	gallery := v1.Group("/gallery")
	gallery.Get("/", galleryHandler.HandleGetAllGallery)
}
