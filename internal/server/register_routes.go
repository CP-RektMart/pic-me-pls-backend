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

	// All, Auth
	authHandler.RegisterLogin(api)
	authHandler.RegisterRegister(api)
	authHandler.RegisterRefreshToken(api)
	authHandler.RegisterLogout(api, authMiddlewares)

	// All, User
	userHandler.RegisterMe(api, authMiddlewares)
	userHandler.RegisterUpdateMe(api, authMiddlewares)

	// Customer, Photographers
	photographersHandler.RegisterGetAllPhotographers(api, authMiddlewares)

	// Photographer, verify
	verifyHandler.RegisterGetCitizenCard(api, authPhotographerMiddlewares)
	verifyHandler.RegisterVerifyCard(api, authPhotographerMiddlewares)
	verifyHandler.RegisterReVerifyCard(api, authPhotographerMiddlewares)

	// Photographer, packages
	packagesHandler.RegisterGetAllPackages(api, authMiddlewares)
	packagesHandler.RegisterCreatePackage(api, authPhotographerMiddlewares)
	packagesHandler.RegisterUpdatePackage(api, authPhotographerMiddlewares)

	// Customer, quotation
	quotationHandler.RegisterAcceptQuotation(api, authMiddlewares)

	// media
	mediaHandler.RegisterCreateMedia(api, authPhotographerMiddlewares)
	mediaHandler.RegisterUpdateMedia(api, authPhotographerMiddlewares)
	mediaHandler.RegisterDeleteMedia(api, authPhotographerMiddlewares)
}
