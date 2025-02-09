package authentication

import (
	"context"

	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/jwt"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
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
}

func NewAuthMiddleware(config *jwt.Config) AuthMiddleware {
	return &authMiddleware{
		config: config,
	}
}

func (r *authMiddleware) Auth(ctx *fiber.Ctx) error {
	tokenByte := ctx.GetReqHeaders()["Authorization"]
	if len(tokenByte) == 0 {
		return ctx.Status(fiber.StatusUnauthorized).JSON(dto.HttpResponse{
			Result: "UNAUTHORIZED",
		})
	}

	if len(tokenByte[0]) < 7 {
		return ctx.Status(fiber.StatusUnauthorized).JSON(dto.HttpResponse{
			Result: "UNAUTHORIZED",
		})
	}

	bearerToken := tokenByte[0][7:]
	if len(bearerToken) == 0 {
		return ctx.Status(fiber.StatusUnauthorized).JSON(dto.HttpResponse{
			Result: "UNAUTHORIZED",
		})
	}

	claims, err := r.validateToken(ctx.UserContext(), bearerToken)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(dto.HttpResponse{
			Result: "UNAUTHORIZED",
		})
	}

	userContext := r.withUserID(ctx.UserContext(), claims.ID)
	ctx.SetUserContext(userContext)

	return ctx.Next()
}

func (r *authMiddleware) AuthAdmin(ctx *fiber.Ctx) error {
	tokenByte := ctx.GetReqHeaders()["Authorization"]
	if len(tokenByte) == 0 {
		return ctx.Status(fiber.StatusUnauthorized).JSON(dto.HttpResponse{
			Result: "UNAUTHORIZED",
		})
	}

	if len(tokenByte[0]) < 7 {
		return ctx.Status(fiber.StatusUnauthorized).JSON(dto.HttpResponse{
			Result: "UNAUTHORIZED",
		})
	}

	bearerToken := tokenByte[0][7:]
	if len(bearerToken) == 0 {
		return ctx.Status(fiber.StatusUnauthorized).JSON(dto.HttpResponse{
			Result: "UNAUTHORIZED",
		})
	}

	claims, err := r.validateToken(ctx.UserContext(), bearerToken)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(dto.HttpResponse{
			Result: "UNAUTHORIZED",
		})
	}

	userContext := r.withUserID(ctx.UserContext(), claims.ID)
	ctx.SetUserContext(userContext)

	if claims.Role != model.UserRoleAdmin {
		return ctx.Status(fiber.StatusUnauthorized).JSON(dto.HttpResponse{
			Result: "UNAUTHORIZED",
		})
	}

	return ctx.Next()
}

func (r *authMiddleware) validateToken(ctx context.Context, bearerToken string) (jwt.JWTentity, error) {
	parsedToken, err := jwt.ParseToken(bearerToken, r.config.AccessTokenSecret)
	if err != nil {
		return jwt.JWTentity{}, errors.Wrap(err, "failed to parse refresh token")
	}

	// TODO: Implement Get Token from Cache
	// cachedToken, err := r.authRepo.GetUserAuthToken(ctx, parsedToken.ID)
	// if err != nil {
	// 	return jwt.JWTentity{}, errors.Wrap(err, "failed to get cached token")
	// }
	_ = ctx
	cachedToken := model.CachedTokens{}

	err = jwt.ValidateToken(cachedToken, parsedToken, false)
	if err != nil {
		return jwt.JWTentity{}, errors.Wrap(err, "failed to validate refresh token")
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
