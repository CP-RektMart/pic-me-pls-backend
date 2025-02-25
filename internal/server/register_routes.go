package server

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/middlewares/authentication"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/services/auth"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/services/category"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/services/citizencard"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/services/media"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/services/message"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/services/object"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/services/packages"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/services/photographers"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/services/quotation"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/services/review"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/services/user"
)

func (s *Server) RegisterRoutes(
	authMiddleware authentication.AuthMiddleware,
	authHandler *auth.Handler,
	userHandler *user.Handler,
	photographersHandler *photographers.Handler,
	citizencardHandler *citizencard.Handler,
	packagesHandler *packages.Handler,
	reviewHandler *review.Handler,
	categoryHandler *category.Handler,
	messageHandler *message.Handler,
	objectHandler *object.Handler,
	quotationHandler *quotation.Handler,
	mediaHandler *media.Handler,
) {
	v1 := s.app.Group("/api/v1")

	// auth
	auth := v1.Group("/auth")
	auth.Post("/login", authHandler.HandleLogin)
	auth.Post("/register", authHandler.HandleRegister)
	auth.Post("/refresh-token", authHandler.HandleRefreshToken)
	auth.Post("/logout", authMiddleware.Auth, authHandler.HandleLogout)

	// all
	{
		all := v1.Group("/", authMiddleware.Auth)

		// me
		me := all.Group("/me")
		me.Get("/", userHandler.HandleGetMe)
		me.Patch("/", userHandler.HandleUpdateMe)

		// photographers
		photographers := all.Group("/photographers")
		photographers.Get("/", photographersHandler.HandleGetAllPhotographers)

		// quotations
		quotations := all.Group("/quotations")
		quotations.Get("/", quotationHandler.HandleListQuotations)
		quotations.Get("/:id", quotationHandler.HandleGetQuotationByID)

		// packages
		packages := all.Group("/packages")
		packages.Get("/", packagesHandler.HandleGetAllPackages)

		// categories
		categories := all.Group("/categories")
		categories.Get("/", categoryHandler.HandleListCategory)
	}

	// customer
	{
		customer := v1.Group("/customer", authMiddleware.AuthCustomer)

		// quotations
		quotations := customer.Group("/quotations")
		quotations.Patch("/:id/confirm", quotationHandler.HandlerConfirmQuotation)
		quotations.Patch("/:id/cancel", quotationHandler.HandlerCancelQuotation)
	}

	// photographer
	{
		photographer := v1.Group("/photographer", authMiddleware.AuthPhotographer)

		// citizen card
		citizencard := photographer.Group("/citizen-card")
		citizencard.Get("/", citizencardHandler.HandleGetCitizenCard)
		citizencard.Post("/verify", citizencardHandler.HandleVerifyCard)
		citizencard.Patch("/reverify", citizencardHandler.HandleReVerifyCard)

		// packages
		packages := photographer.Group("/packages")
		packages.Post("/", packagesHandler.HandleCreatePackage)
		packages.Patch("/:packageId", packagesHandler.HandleUpdatePackage)

		// media
		media := photographer.Group("/media")
		media.Post("/", mediaHandler.HandleCreateMedia)
		media.Patch("/:mediaId", mediaHandler.HandleUpdateMedia)
		media.Delete("/:mediaId", mediaHandler.HandleDeleteMedia)
	}

	// admin
	{
		admin := v1.Group("/admin", authMiddleware.AuthAdmin)

		// category
		categories := admin.Group("/categories")
		categories.Post("/", categoryHandler.HandleCreateCategory)
		categories.Patch("/:id", categoryHandler.HandleUpdateCategory)
		categories.Delete("/:id", categoryHandler.HandleDeleteCategory)
	}
}
