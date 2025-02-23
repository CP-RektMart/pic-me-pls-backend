package server

import (
	"context"

	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
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
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humafiber"
)

func (s *Server) RegisterRoutes(
	authMiddleware authentication.AuthMiddleware,
	authHandler *auth.Handler,
	userHandler *user.Handler,
	verifyHandler *verify.Handler,
	photographersHandler *photographers.Handler,
	packagesHandler *packages.Handler,
	reviewHandler *review.Handler,
	categoryHandler *category.Handler,
	messageHandler *message.Handler,
	objectHandler *object.Handler,
	quotationHandler *quotation.Handler,
	mediaHandler *media.Handler,
) {
	config := huma.DefaultConfig("Pic-Me-Pls Backend", "0.0.1")
	config.Components.SecuritySchemes = map[string]*huma.SecurityScheme{
		"bearer": {
			Type:         "http",
			Scheme:       "bearer",
			BearerFormat: "JWT",
		},
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

	// Middlewares
	authMiddlewares := huma.Middlewares{
		func(ctx huma.Context, next func(huma.Context)) {
			authMiddleware.Auth(ctx, next, api)
		},
	}
	authPhotographerMiddlewares := huma.Middlewares{
		func(ctx huma.Context, next func(huma.Context)) {
			authMiddleware.AuthPhotographer(ctx, next, api)
		},
	}
	// authAdminMiddlewares := huma.Middlewares{
	// 	func(ctx huma.Context, next func(huma.Context)) {
	// 		authMiddleware.AuthAdmin(ctx, next, api)
	// 	},
	// }

	// auth
	authHandler.RegisterLogin(api)
	authHandler.RegisterRegister(api)
	authHandler.RegisterRefreshToken(api)
	authHandler.RegisterLogout(api, authMiddlewares)

	// user
	userHandler.RegisterMe(api, authMiddlewares)
	userHandler.RegisterUpdateMe(api, authMiddlewares)

	// Photographer, verify
	verifyHandler.RegisterGetCitizenCard(api, authPhotographerMiddlewares)
	verifyHandler.RegisterVerifyCard(api, authPhotographerMiddlewares)
	verifyHandler.RegisterReVerifyCard(api, authPhotographerMiddlewares)

	// Photographers
	photographersHandler.RegisterGetAllPhotographers(api, authMiddlewares)

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
			o.Middlewares = append(o.Middlewares, func(ctx huma.Context, next func(huma.Context)) {
				authMiddleware.AuthPhotographer(ctx, next, packages)
			})
		})
		huma.Patch(packages, basePath+"/{packageId}", packagesHandler.HandleUpdatePackage, func(o *huma.Operation) {
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
			o.Middlewares = append(o.Middlewares, func(ctx huma.Context, next func(huma.Context)) {
				authMiddleware.Auth(ctx, next, quotations)
			})
		})
	}

	// media
	{
		basePath := "/api/v1/media"
		media := humafiber.New(s.app, config)

		huma.Post(media, basePath, mediaHandler.HandleCreateMedia, func(o *huma.Operation) {
			o.Middlewares = append(o.Middlewares, func(ctx huma.Context, next func(huma.Context)) {
				authMiddleware.AuthPhotographer(ctx, next, media)
			})
		})
		huma.Patch(media, basePath+"/{mediaId}", mediaHandler.HandleUpdateMedia, func(o *huma.Operation) {
			o.Middlewares = append(o.Middlewares, func(ctx huma.Context, next func(huma.Context)) {
				authMiddleware.AuthPhotographer(ctx, next, media)
			})
		})
		huma.Delete(media, basePath+"/{mediaId}", mediaHandler.HandleDeleteMedia, func(o *huma.Operation) {
			o.Middlewares = append(o.Middlewares, func(ctx huma.Context, next func(huma.Context)) {
				authMiddleware.AuthPhotographer(ctx, next, media)
			})
		})
	}
}
