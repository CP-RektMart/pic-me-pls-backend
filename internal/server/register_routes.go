package server

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/middlewares/authentication"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/services/auth"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/services/category"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/services/media"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/services/message"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/services/object"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/services/packages"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/services/photographers"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/services/quotation"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/services/review"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/services/user"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/services/verify"
)

func (s *Server) RegisterRoutes(
	authMiddleware authentication.AuthMiddleware,
	authHandler *auth.Handler,
	userHandler *user.Handler,
	photographersHandler *photographers.Handler,
	verifyHandler *verify.Handler,
	packagesHandler *packages.Handler,
	reviewHandler *review.Handler,
	categoryHandler *category.Handler,
	messageHandler *message.Handler,
	objectHandler *object.Handler,
	quotationHandler *quotation.Handler,
	mediaHandler *media.Handler,
) {
	v1 := s.app.Group("/api/v1")

	// all, auth
	auth := v1.Group("/auth")
	auth.Post("/login", authHandler.HandleLogin)
	auth.Post("/register", authHandler.HandleRegister)
	auth.Post("/refresh-token", authHandler.HandleRefreshToken)
	auth.Post("/logout", authMiddleware.Auth, authHandler.HandleLogout)

	// all, user
	v1.Get("/me", authMiddleware.Auth, userHandler.HandleGetMe)
	v1.Patch("/me", authMiddleware.Auth, userHandler.HandleUpdateMe)

	// verify citizen card
	photographer := v1.Group("/photographer")
	photographer.Get("/citizen-card", authMiddleware.AuthPhotographer, verifyHandler.HandleGetCitizenCard)
	photographer.Post("/verify", authMiddleware.AuthPhotographer, verifyHandler.HandleVerifyCard)
	photographer.Patch("/reverify", authMiddleware.AuthPhotographer, verifyHandler.HandleReVerifyCard)

	// get photographer
	photographers := v1.Group("/photographers")
	photographers.Get("/", photographersHandler.HandleGetAllPhotographers)

	auth.Post("/logout", authMiddleware.Auth, authHandler.HandleLogout)

	// package
	packages := v1.Group("/packages")
	packages.Get("/", packagesHandler.HandleGetAllPackages)
	packages.Post("/", authMiddleware.AuthPhotographer, packagesHandler.HandleCreatePackage)
	packages.Patch("/:packageId", authMiddleware.AuthPhotographer, packagesHandler.HandleUpdatePackage)

	// quotation
	quotation := v1.Group("/quotations")
	quotation.Get("/", authMiddleware.Auth, quotationHandler.HandleListQuotations)
	quotation.Get("/:id", authMiddleware.Auth, quotationHandler.HandleGetQuotationByID)
	quotation.Patch("/:id/confirm", authMiddleware.Auth, quotationHandler.HandlerConfirmQuotation)
	quotation.Patch("/:id/cancel", authMiddleware.Auth, quotationHandler.HandlerCancelQuotation)

	// category
	category := v1.Group("/categories")
	category.Post("/", authMiddleware.AuthAdmin, categoryHandler.HandleCreateCategory)
	category.Patch("/:id", authMiddleware.AuthAdmin, categoryHandler.HandleUpdateCategory)
	category.Get("/", categoryHandler.HandleListCategory)
	category.Delete("/:id", authMiddleware.AuthAdmin, categoryHandler.HandleDeleteCategory)

	// media
	media := v1.Group("/media")
	media.Post("/", authMiddleware.AuthPhotographer, mediaHandler.HandleCreateMedia)
	media.Patch("/:mediaId", authMiddleware.AuthPhotographer, mediaHandler.HandleUpdateMedia)
	media.Delete("/:mediaId", authMiddleware.AuthPhotographer, mediaHandler.HandleDeleteMedia)
}
