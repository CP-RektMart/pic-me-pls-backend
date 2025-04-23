package server

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/middlewares/authentication"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/services/admin"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/services/auth"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/services/category"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/services/citizencard"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/services/customer"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/services/media"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/services/message"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/services/objects"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/services/packages"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/services/photographers"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/services/quotation"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/services/report"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/services/review"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/services/stripe"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/services/user"
	"github.com/gofiber/contrib/websocket"
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
	objectsHandler *objects.Handler,
	quotationHandler *quotation.Handler,
	mediaHandler *media.Handler,
	customerHandler *customer.Handler,
	messageHandler *message.Handler,
	stripeHandler *stripe.Handler,
	adminHandler *admin.Handler,
	reportHandler *report.Handler,
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
		all := v1.Group("/")

		// customer
		customer := all.Group("/customers")
		customer.Get("/:id", customerHandler.HandlerCustomerPublicProfile)

		// objects
		objects := all.Group("/objects")
		objects.Post("/", objectsHandler.Upload)
		objects.Delete("/", objectsHandler.Delete)

		// me
		me := all.Group("/me")
		me.Get("/", authMiddleware.Auth, userHandler.HandleGetMe)
		me.Patch("/", authMiddleware.Auth, userHandler.HandleUpdateMe)

		// photographers
		photographers := all.Group("/photographers")
		photographers.Get("/", photographersHandler.HandleGetAllPhotographers)
		photographers.Get("/:id", photographersHandler.HandlerGetPhotographerByID)

		// quotations
		quotations := all.Group("/quotations")
		quotations.Get("/", authMiddleware.Auth, quotationHandler.HandleListQuotations)
		quotations.Get("/:id", authMiddleware.Auth, quotationHandler.HandleGetQuotationByID)

		// packages
		packages := all.Group("/packages")
		packages.Get("/", packagesHandler.HandleGetAllPackages)
		packages.Get("/:id", packagesHandler.HandleGetPackageByID)
		packages.Get("/:packageId/reviews", packagesHandler.HandleGetPackageReviews)

		// categories
		categories := all.Group("/categories")
		categories.Get("/", categoryHandler.HandleListCategory)

		// messages
		message := all.Group("/messages")
		message.Use("/ws", messageHandler.HandleSupportWebAPI, authMiddleware.Auth, messageHandler.HandleWebsocket)
		message.Get("/ws", websocket.New(messageHandler.HandleRealTimeMessages))
		message.Get("/", authMiddleware.Auth, messageHandler.HandleListMessages)

	}

	// customer
	{
		customer := v1.Group("/customer", authMiddleware.AuthCustomer)

		// quotations
		quotations := customer.Group("/quotations")
		quotations.Patch("/:id/confirm", quotationHandler.HandlerConfirmQuotation)
		quotations.Patch("/:id/cancel", quotationHandler.HandlerCancelQuotation)
		quotations.Patch("/:id/complete", quotationHandler.HandleCompleteQuotation)
		quotations.Post("/:quotationId/review", reviewHandler.HandleCreateReview)
		quotations.Patch("/:quotationId/review/:id", reviewHandler.HandleUpdateReview)
		quotations.Delete("/:quotationId/review/:id", reviewHandler.HandleDeleteReview)

		// reports
		reports := customer.Group("/reports")
		reports.Post("/", reportHandler.HandleCreateReport)
		reports.Get("/", reportHandler.HandleGetAllReports)
		reports.Get("/:id", reportHandler.HandleGetReportByID)
		reports.Patch("/:id", reportHandler.HandleUpdateReport)
	}

	// photographer
	{
		photographer := v1.Group("/photographer", authMiddleware.AuthPhotographer)

		me := photographer.Group("/me")
		me.Get("/", photographersHandler.HandleGetMe)

		// citizen card
		citizencard := photographer.Group("/citizen-card")
		citizencard.Get("/", citizencardHandler.HandleGetCitizenCard)
		citizencard.Post("/verify", citizencardHandler.HandleVerifyCard)
		citizencard.Patch("/reverify", citizencardHandler.HandleReVerifyCard)

		// packages
		packages := photographer.Group("/packages")
		packages.Post("/", packagesHandler.HandleCreatePackage)
		packages.Patch("/:id", packagesHandler.HandleUpdatePackage)
		packages.Get("/", packagesHandler.HandlerListPhotographerPackages)
		packages.Delete("/:id", packagesHandler.HandleDeletePackage)

		// media
		media := photographer.Group("/media")
		media.Post("/", mediaHandler.HandleCreateMedia)
		media.Patch("/:mediaId", mediaHandler.HandleUpdateMedia)
		media.Delete("/:mediaId", mediaHandler.HandleDeleteMedia)

		// quotations
		quotations := photographer.Group("/quotations")
		quotations.Post("/", quotationHandler.HandleCreateQuotation)
		quotations.Patch("/:id", quotationHandler.HandleUpdateQuotation)
		quotations.Post(":id/preview", quotationHandler.HandleCreatePreviewPhoto)
	}

	// admin
	{
		admin := v1.Group("/admin", authMiddleware.AuthAdmin)

		// category
		categories := admin.Group("/categories")
		categories.Post("/", categoryHandler.HandleCreateCategory)
		categories.Patch("/:id", categoryHandler.HandleUpdateCategory)
		categories.Delete("/:id", categoryHandler.HandleDeleteCategory)

		// photographers
		photographers := admin.Group("/photographers")
		photographers.Get("/", adminHandler.HandleListPhotographers)
		photographers.Get("/:photographerID", adminHandler.HandleGetPhotographerByID)
		photographers.Patch("/:photographerID/verify", adminHandler.HandleVerifyPhotographer)

		// users
		users := admin.Group("/users")
		users.Get("/", adminHandler.HandleGetAllUsers)
		users.Get("/:id", adminHandler.HandleGetUserByID)
		users.Patch("/:userID/role", adminHandler.HandleAssignAdmin)

		// citizendCard

		// photographer
		photographer := admin.Group("/photographer")
		photographer.Patch("/:id/ban", adminHandler.HandleBanPhotographer)
		photographer.Patch("/:id/unban", adminHandler.HandleUnbanPhotographer)

		// packages
		packages := admin.Group("/packages")
		packages.Delete("/:packageID", adminHandler.HandleDeletePackageByID)

		//reports
		reports := admin.Group("/reports")
		reports.Get("/", reportHandler.HandleAdminGetAllReports)
		reports.Patch("/:reportID/accept", adminHandler.HandleAcceptReport)
		reports.Patch("/:reportID/decline", adminHandler.HandleDeclineReport)
	}

	// stripe
	{
		stripe := v1.Group("/stripe")
		stripe.Post("/checkout/quotations/:id", authMiddleware.AuthCustomer, stripeHandler.HandleCreateCheckoutSession)
		stripe.Post("/webhook", stripeHandler.HandleStripeWebhook)
	}
}
