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
		photographers := humafiber.New(s.app, config)
		photographers.UseMiddleware(func(ctx huma.Context, next func(huma.Context)) {
			authMiddleware.AuthPhotographer(ctx, next, photographers)
		})

		// Citizen Card
		huma.Get(photographers, basePath+"/citizen-card", photographerHandler.HandleGetCitizenCard)
		huma.Post(photographers, basePath+"/verify", photographerHandler.HandleVerifyCard)
		huma.Patch(photographers, basePath+"/reverify", photographerHandler.HandleReVerifyCard)
	}

	// Photographer
	{
		basePath := "/api/v1/photographers"
		photographers := humafiber.New(s.app, config)
		photographers.UseMiddleware(func(ctx huma.Context, next func(huma.Context)) {
			authMiddleware.Auth(ctx, next, photographers)
		})

		// Photographers
		huma.Get(photographers, basePath, photographerHandler.HandleGetAllPhotographers)
	}

	// package
	{

		basePath := "/api/v1/packages"
		packages := humafiber.New(s.app, config)

		huma.Get(packages, basePath, packagesHandler.HandleGetAllPackages, func(o *huma.Operation) {
			o.Middlewares = append(o.Middlewares, func(ctx huma.Context, next func(huma.Context)) {
				authMiddleware.Auth(ctx, next, packages)
			})
		})
		huma.Post(packages, basePath, packagesHandler.HandleCreatePackage, func(o *huma.Operation) {
			o.DefaultStatus = 201
			o.Middlewares = append(o.Middlewares, func(ctx huma.Context, next func(huma.Context)) {
				authMiddleware.AuthPhotographer(ctx, next, packages)
			})
		})
		huma.Patch(packages, basePath+"/{packageId}", packagesHandler.HandleUpdatePackage, func(o *huma.Operation) {
			o.DefaultStatus = 204
			o.Middlewares = append(o.Middlewares, func(ctx huma.Context, next func(huma.Context)) {
				authMiddleware.AuthPhotographer(ctx, next, packages)
			})
		})
	}

	// quotation
	{
		basePath := "/api/v1/quotations"
		quotations := humafiber.New(s.app, config)

		huma.Patch(quotations, basePath+"/{id}/accept", quotationHandler.Accept, func(o *huma.Operation) {
			o.DefaultStatus = 204
			o.Middlewares = append(o.Middlewares, func(ctx huma.Context, next func(huma.Context)) {
				authMiddleware.Auth(ctx, next, quotations)
			})
		})
	}
}
