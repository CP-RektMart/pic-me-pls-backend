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
	// health check
	huma.Get(s.Api, "/health", func(ctx context.Context, input *struct{}) (*dto.HumaHttpResponse[string], error) {
		return &dto.HumaHttpResponse[string]{
			Body: dto.HttpResponse[string]{
				Result: "ok",
			},
		}, nil
	})

	// All, Auth
	authHandler.RegisterLogin(s.Api)
	authHandler.RegisterRegister(s.Api)
	authHandler.RegisterRefreshToken(s.Api)
	authHandler.RegisterLogout(s.Api, authMiddleware.Auth)

	// All, User
	userHandler.RegisterMe(s.Api, authMiddleware.Auth)
	userHandler.RegisterUpdateMe(s.Api, authMiddleware.Auth)

	// Customer, Photographers
	photographersHandler.RegisterGetAllPhotographers(s.Api, authMiddleware.Auth)

	// Customer, quotation
	quotationHandler.RegisterAcceptQuotation(s.Api, authMiddleware.Auth)

	// Photographer, verify
	verifyHandler.RegisterGetCitizenCard(s.Api, authMiddleware.AuthPhotographer)
	verifyHandler.RegisterVerifyCard(s.Api, authMiddleware.AuthPhotographer)
	verifyHandler.RegisterReVerifyCard(s.Api, authMiddleware.AuthPhotographer)

	// Photographer, packages
	packagesHandler.RegisterGetAllPackages(s.Api, authMiddleware.Auth)
	packagesHandler.RegisterCreatePackage(s.Api, authMiddleware.AuthPhotographer)
	packagesHandler.RegisterUpdatePackage(s.Api, authMiddleware.AuthPhotographer)

	// Photographer, media
	mediaHandler.RegisterCreateMedia(s.Api, authMiddleware.AuthPhotographer)
	mediaHandler.RegisterUpdateMedia(s.Api, authMiddleware.AuthPhotographer)
	mediaHandler.RegisterDeleteMedia(s.Api, authMiddleware.AuthPhotographer)
}
