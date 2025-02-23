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
	config.Components.SecuritySchemes = map[string]*huma.SecurityScheme{
		"bearer": {
			Type:         "http",
			Scheme:       "bearer",
			BearerFormat: "JWT",
		},
	}
	config.Security = []map[string][]string{
		{"bearer": nil},
	}
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

		huma.Post(auth, basePath+"/login", authHandler.HandleLogin, func(o *huma.Operation) {
			o.Security = []map[string][]string{}
		})
		huma.Post(auth, basePath+"/register", authHandler.HandleRegister, func(o *huma.Operation) {
			o.Security = []map[string][]string{}
		})
		huma.Post(auth, basePath+"/refresh-token", authHandler.HandleRefreshToken, func(o *huma.Operation) {
			o.Security = []map[string][]string{}
		})
		huma.Post(auth, basePath+"/logout", authHandler.HandleLogout, func(o *huma.Operation) {
			o.DefaultStatus = 204
			o.Middlewares = append(o.Middlewares, func(ctx huma.Context, next func(huma.Context)) {
				authMiddleware.Auth(ctx, next, auth)
			})
		})
	}

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

	// Citizen Card
	{
		basePath := "/api/v1/photographers"
		photographer := humafiber.New(s.app, config)
		photographer.UseMiddleware(func(ctx huma.Context, next func(huma.Context)) {
			authMiddleware.AuthPhotographer(ctx, next, photographer)
		})

		// Citizen Card
		huma.Get(photographer, basePath+"/citizen-card", photographerHandler.HandleGetCitizenCard)
		huma.Post(photographer, basePath+"/verify", photographerHandler.HandleVerifyCard)
		huma.Patch(photographer, basePath+"/reverify", photographerHandler.HandleReVerifyCard)
	}

	// Photographer
	{
		basePath := "/api/v1/photographers"
		photographer := humafiber.New(s.app, config)
		photographer.UseMiddleware(func(ctx huma.Context, next func(huma.Context)) {
			authMiddleware.Auth(ctx, next, photographer)
		})

		// Photographers
		huma.Get(photographer, basePath, photographerHandler.HandleGetAllPhotographers)
	}

	// // package
	// packages := v1.Group("/packages")
	// packages.Get("/", packagesHandler.HandleGetAllPackages)
	// packages.Post("/", authMiddleware.AuthPhotographer, packagesHandler.HandleCreatePackage)
	// packages.Patch("/:packageId", authMiddleware.AuthPhotographer, packagesHandler.HandleUpdatePackage)

	// // quotation
	// quotation := v1.Group("/quotations")
	// quotation.Patch("/:id/accept", authMiddleware.Auth, quotationHandler.Accept)
}
