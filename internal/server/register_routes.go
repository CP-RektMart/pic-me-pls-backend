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
		auth := humafiber.New(s.app, config)

		huma.Post(auth, basePath+"/login", authHandler.HandleLogin)
		huma.Post(auth, basePath+"/register", authHandler.HandleRegister)
		huma.Post(auth, basePath+"/refresh-token", authHandler.HandleRefreshToken)
		huma.Post(auth, basePath+"/logout", authHandler.HandleLogout, func(o *huma.Operation) {
			o.DefaultStatus = 204
			o.Middlewares = append(o.Middlewares, func(ctx huma.Context, next func(huma.Context)) {
				authMiddleware.Auth(ctx, next, auth)
			})
		})
	}

	// v1 := s.app.Group("/api/v1")

	// user
	{
		basePath := "/api/v1/me"
		user := humafiber.New(s.app, config)
		user.UseMiddleware(func(ctx huma.Context, next func(huma.Context)) {
			authMiddleware.Auth(ctx, next, user)
		})

		huma.Get(user, basePath, userHandler.HandleGetMe)
		huma.Patch(user, basePath, userHandler.HandleUpdateMe)
	}

	// // verify citizen card
	// photographer := v1.Group("/photographer")
	// photographer.Get("/citizen-card", authMiddleware.AuthPhotographer, photographerHandler.HandleGetCitizenCard)
	// photographer.Post("/verify", authMiddleware.AuthPhotographer, photographerHandler.HandleVerifyCard)
	// photographer.Patch("/reverify", authMiddleware.AuthPhotographer, photographerHandler.HandleReVerifyCard)

	// // get photographer
	// photographers := v1.Group("/photographers")
	// photographers.Get("/", photographerHandler.HandleGetAllPhotographers)

	// // package
	// packages := v1.Group("/packages")
	// packages.Get("/", packagesHandler.HandleGetAllPackages)
	// packages.Post("/", authMiddleware.AuthPhotographer, packagesHandler.HandleCreatePackage)
	// packages.Patch("/:packageId", authMiddleware.AuthPhotographer, packagesHandler.HandleUpdatePackage)

	// // quotation
	// quotation := v1.Group("/quotations")
	// quotation.Patch("/:id/accept", authMiddleware.Auth, quotationHandler.Accept)
}
