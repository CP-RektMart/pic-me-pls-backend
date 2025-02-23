package authentication

import (
	"context"
	"net/http"

	"github.com/CP-RektMart/pic-me-pls-backend/internal/jwt"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/logger"
	"github.com/cockroachdb/errors"
	"github.com/danielgtaylor/huma/v2"
)

var (
	ErrInvalidToken = errors.New("INVALID_TOKEN")
	ErrNoPermission = errors.New("NO_PERMISSION")
)

type AuthMiddleware interface {
	Auth(ctx huma.Context, next func(huma.Context), api huma.API)
	AuthAdmin(ctx huma.Context, next func(huma.Context), api huma.API)
	AuthPhotographer(ctx huma.Context, next func(huma.Context), api huma.API)
	GetUserIDFromContext(ctx context.Context) (uint, error)
	GetJWTEntityFromContext(ctx context.Context) (jwt.JWTentity, error)
}

type authMiddleware struct {
	jwtService *jwt.JWT
}

func NewAuthMiddleware(jwtService *jwt.JWT) AuthMiddleware {
	return &authMiddleware{
		jwtService: jwtService,
	}
}

func (r *authMiddleware) Auth(ctx huma.Context, next func(huma.Context), api huma.API) {
	tokenByte := ctx.Header("Authorization")

	if len(tokenByte) < 7 {
		if err := huma.WriteErr(api, ctx, http.StatusUnauthorized, "invalid token", ErrInvalidToken); err != nil {
			logger.ErrorContext(ctx.Context(), "failed to write error", err)
		}
		return
	}

	bearerToken := tokenByte[7:]
	if len(bearerToken) == 0 {
		if err := huma.WriteErr(api, ctx, http.StatusUnauthorized, "invalid token", ErrInvalidToken); err != nil {
			logger.ErrorContext(ctx.Context(), "failed to write error", err)
		}
		return
	}

	claims, err := r.validateToken(ctx.Context(), bearerToken)
	if err != nil {
		if err := huma.WriteErr(api, ctx, http.StatusUnauthorized, "invalid token", ErrInvalidToken); err != nil {
			logger.ErrorContext(ctx.Context(), "failed to write error", err)
		}
		return
	}

	userContext := r.withJWTEntity(ctx, claims)
	next(userContext)
}

func (r *authMiddleware) AuthAdmin(ctx huma.Context, next func(huma.Context), api huma.API) {
	tokenByte := ctx.Header("Authorization")
	if len(tokenByte) < 7 {
		if err := huma.WriteErr(api, ctx, http.StatusUnauthorized, "invalid token", ErrInvalidToken); err != nil {
			logger.ErrorContext(ctx.Context(), "failed to write error", err)
		}
		return
	}

	bearerToken := tokenByte[7:]
	if len(bearerToken) == 0 {
		if err := huma.WriteErr(api, ctx, http.StatusUnauthorized, "invalid token", ErrInvalidToken); err != nil {
			logger.ErrorContext(ctx.Context(), "failed to write error", err)
		}
		return
	}

	claims, err := r.validateToken(ctx.Context(), bearerToken)
	if err != nil {
		if err := huma.WriteErr(api, ctx, http.StatusUnauthorized, "invalid token", ErrInvalidToken); err != nil {
			logger.ErrorContext(ctx.Context(), "failed to write error", err)
		}
		return
	}

	if claims.Role != model.UserRoleAdmin {
		if err := huma.WriteErr(api, ctx, http.StatusForbidden, "no permission", ErrNoPermission); err != nil {
			logger.ErrorContext(ctx.Context(), "failed to write error", err)
		}
		return
	}

	userContext := r.withJWTEntity(ctx, claims)
	next(userContext)
}

func (r *authMiddleware) AuthPhotographer(ctx huma.Context, next func(huma.Context), api huma.API) {
	tokenByte := ctx.Header("Authorization")

	if len(tokenByte) < 7 {
		if err := huma.WriteErr(api, ctx, http.StatusUnauthorized, "invalid token", ErrInvalidToken); err != nil {
			logger.ErrorContext(ctx.Context(), "failed to write error", err)
		}
		return
	}

	bearerToken := tokenByte[7:]
	if len(bearerToken) == 0 {
		if err := huma.WriteErr(api, ctx, http.StatusUnauthorized, "invalid token", ErrInvalidToken); err != nil {
			logger.ErrorContext(ctx.Context(), "failed to write error", err)
		}
		return
	}

	claims, err := r.validateToken(ctx.Context(), bearerToken)
	if err != nil {
		if err := huma.WriteErr(api, ctx, http.StatusUnauthorized, "invalid token", ErrInvalidToken); err != nil {
			logger.ErrorContext(ctx.Context(), "failed to write error", err)
		}
		return
	}

	if claims.Role != model.UserRolePhotographer {
		if err := huma.WriteErr(api, ctx, http.StatusForbidden, "no permission", ErrNoPermission); err != nil {
			logger.ErrorContext(ctx.Context(), "failed to write error", err)
		}
		return
	}

	userContext := r.withJWTEntity(ctx, claims)
	next(userContext)
}

func (r *authMiddleware) validateToken(ctx context.Context, bearerToken string) (jwt.JWTentity, error) {
	parsedToken, err := r.jwtService.ParseToken(bearerToken, false)
	if err != nil {
		return jwt.JWTentity{}, errors.Wrap(err, "failed to parse  token")
	}

	cachedToken, err := r.jwtService.GetCachedTokens(ctx, parsedToken.ID)
	if err != nil {
		return jwt.JWTentity{}, errors.Wrap(err, "failed to get cached token")
	}

	err = r.jwtService.ValidateToken(*cachedToken, parsedToken, false)
	if err != nil {
		return jwt.JWTentity{}, errors.Wrap(err, "failed to validate token")
	}

	return parsedToken, nil

}

type jwtEntityContext struct{}

func (r *authMiddleware) withJWTEntity(ctx huma.Context, jwtEntity jwt.JWTentity) huma.Context {
	return huma.WithValue(ctx, jwtEntityContext{}, jwtEntity)
}

func (r *authMiddleware) GetJWTEntityFromContext(ctx context.Context) (jwt.JWTentity, error) {
	jwtEntity, ok := ctx.Value(jwtEntityContext{}).(jwt.JWTentity)

	if !ok {
		return jwt.JWTentity{}, errors.New("failed to get jwt entity from context")
	}

	return jwtEntity, nil
}

func (r *authMiddleware) GetUserIDFromContext(ctx context.Context) (uint, error) {
	jwtEntity, err := r.GetJWTEntityFromContext(ctx)
	if err != nil {
		return 0, errors.Wrap(err, "failed to get jwt entity from context")
	}

	return jwtEntity.ID, nil
}
