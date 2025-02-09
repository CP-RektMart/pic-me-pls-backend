package authentication

import (
	"context"
	"fmt"

	"github.com/CP-RektMart/pic-me-pls-backend/internal/jwt"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

var (
	ErrInvalidToken = errors.New("INVALID_TOKEN")
)

type AuthMiddleware interface {
	Auth(ctx *fiber.Ctx) error
	AuthAdmin(ctx *fiber.Ctx) error
	GetUserIDFromContext(ctx context.Context) (uint, error)
}

type authMiddleware struct {
	config *jwt.Config
	cache  *redis.Client
}

func NewAuthMiddleware(config *jwt.Config, cache *redis.Client) AuthMiddleware {
	return &authMiddleware{
		config: config,
		cache: cache,
	}
}

func (r *authMiddleware) Auth(ctx *fiber.Ctx) error {
	tokenByte := ctx.GetReqHeaders()["Authorization"]
	if len(tokenByte) == 0 {
		return apperror.UnAuthorized("UNAUTHORIZED", fmt.Errorf("no header"))
	}

	if len(tokenByte[0]) < 7 {
		return apperror.UnAuthorized("UNAUTHORIZED", fmt.Errorf("invalid header"))
	}

	bearerToken := tokenByte[0][7:]
	if len(bearerToken) == 0 {
		return apperror.UnAuthorized("UNAUTHORIZED", fmt.Errorf("no bearer keyword"))
	}

	claims, err := r.validateToken(ctx.UserContext(), bearerToken)
	if err != nil {
		return apperror.UnAuthorized("UNAUTHORIZED", errors.Wrap(err, "failed to validate token"))
	}

	userContext := r.withUserID(ctx.UserContext(), claims.ID)
	ctx.SetUserContext(userContext)

	return ctx.Next()
}

func (r *authMiddleware) AuthAdmin(ctx *fiber.Ctx) error {
	tokenByte := ctx.GetReqHeaders()["Authorization"]
	if len(tokenByte) == 0 {
		return apperror.UnAuthorized("UNAUTHORIZED", fmt.Errorf("no header"))
	}

	if len(tokenByte[0]) < 7 {
		return apperror.UnAuthorized("UNAUTHORIZED", fmt.Errorf("invalid header"))
	}

	bearerToken := tokenByte[0][7:]
	if len(bearerToken) == 0 {
		return apperror.UnAuthorized("UNAUTHORIZED", fmt.Errorf("no bearer keyword"))
	}

	claims, err := r.validateToken(ctx.UserContext(), bearerToken)
	if err != nil {
		return apperror.UnAuthorized("UNAUTHORIZED", errors.Wrap(err, "failed to validate token"))
	}

	userContext := r.withUserID(ctx.UserContext(), claims.ID)
	ctx.SetUserContext(userContext)

	if claims.Role != model.UserRoleAdmin {
		return apperror.Forbidden("FORBIDDEN", fmt.Errorf("user is not admin"))
	}

	return ctx.Next()
}

func (r *authMiddleware) validateToken(ctx context.Context, bearerToken string) (jwt.JWTentity, error) {
	parsedToken, err := jwt.ParseToken(bearerToken, r.config.AccessTokenSecret)
	if err != nil {
		return jwt.JWTentity{}, errors.Wrap(err, "failed to parse  token")
	}

	cachedToken, err := jwt.GetCachedTokens(ctx, r.cache, parsedToken.ID)
	if err != nil {
		return jwt.JWTentity{}, errors.Wrap(err, "failed to get cached token")
	}

	err = jwt.ValidateToken(*cachedToken, parsedToken, false)
	if err != nil {
		return jwt.JWTentity{}, errors.Wrap(err, "failed to validate token")
	}

	return parsedToken, nil

}

type userIDContext struct{}

func (r *authMiddleware) withUserID(ctx context.Context, userID uint) context.Context {
	return context.WithValue(ctx, userIDContext{}, userID)
}

func (r *authMiddleware) GetUserIDFromContext(ctx context.Context) (uint, error) {
	userID, ok := ctx.Value(userIDContext{}).(uint)

	if !ok {
		return 0, errors.New("failed to get user id from context")
	}

	return userID, nil
}
