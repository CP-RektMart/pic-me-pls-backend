package server

import (
	"context"

	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/middlewares/authentication"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/services/auth"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/services/category"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/services/message"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/services/object"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/services/packages"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/services/photographer"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/services/quotation"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/services/review"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/services/user"
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humafiber"
)

type ReviewInput struct {
	Body struct {
		Author  string `json:"author" maxLength:"10" doc:"Author of the review"`
		Rating  int    `json:"rating" minimum:"1" maximum:"5" doc:"Rating from 1 to 5"`
		Message string `json:"message,omitempty" maxLength:"100" doc:"Review message"`
	}
}

func (s *Server) RegisterRoutes(
	authMiddleware authentication.AuthMiddleware,
	authHandler *auth.Handler,
	userHandler *user.Handler,
	photographerHandler *photographer.Handler,
	packagesHandler *packages.Handler,
	reviewHandler *review.Handler,
	categoryHandler *category.Handler,
	messageHandler *message.Handler,
	objectHandler *object.Handler,
	quotationHandler *quotation.Handler,
) {
	config := huma.DefaultConfig("Pic-Me-Pls Backend", "0.0.1")
	api := humafiber.New(s.app, config)

	// health check
	{
		huma.Get(api, "/health", func(ctx context.Context, input *struct{}) (*dto.HumaHttpResponse[string], error) {
			return &dto.HumaHttpResponse[string]{
				Body: dto.HttpResponse[string]{
					Result: "ok",
				},
			}, nil
		})
	}

	// auth
	{
		basePath := "/api/v1/auth"
		huma.Post(api, basePath+"/login", authHandler.HandleLogin)
		// auth.Post("/login", authHandler.HandleLogin)
		// auth.Post("/register", authHandler.HandleRegister)
		// auth.Post("/refresh-token", authHandler.HandleRefreshToken)
		// auth.Post("/logout", authMiddleware.Auth, authHandler.HandleLogout)
	}

	v1 := s.app.Group("/api/v1")

	// user
	v1.Get("/me", authMiddleware.Auth, userHandler.HandleGetMe)
	v1.Patch("/me", authMiddleware.Auth, userHandler.HandleUpdateMe)

	// verify citizen card
	photographer := v1.Group("/photographer")
	photographer.Get("/citizen-card", authMiddleware.AuthPhotographer, photographerHandler.HandleGetCitizenCard)
	photographer.Post("/verify", authMiddleware.AuthPhotographer, photographerHandler.HandleVerifyCard)
	photographer.Patch("/reverify", authMiddleware.AuthPhotographer, photographerHandler.HandleReVerifyCard)

	// get photographer
	photographers := v1.Group("/photographers")
	photographers.Get("/", photographerHandler.HandleGetAllPhotographers)

	// package
	packages := v1.Group("/packages")
	packages.Get("/", packagesHandler.HandleGetAllPackages)
	packages.Post("/", authMiddleware.AuthPhotographer, packagesHandler.HandleCreatePackage)
	packages.Patch("/:packageId", authMiddleware.AuthPhotographer, packagesHandler.HandleUpdatePackage)

	// quotation
	quotation := v1.Group("/quotations")
	quotation.Patch("/:id/accept", authMiddleware.Auth, quotationHandler.Accept)
}
