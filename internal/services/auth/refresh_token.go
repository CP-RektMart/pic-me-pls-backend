package auth

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// @Summary			Refresh Token
// @Description		Refresh Token
// @Tags			auth
// @Router			/api/v1/auth/refresh-token [POST]
// @Success			200 {object}	dto.HttpResponse{result=dto.TokenResponse}
// @Failure			400	{object}	dto.HttpResponse
// @Failure			500	{object}	dto.HttpResponse
func (h *Handler) HandleRefreshToken(c *fiber.Ctx) error {
	ctx := c.UserContext()

	var req dto.RefreshTokenRequest
	if err := c.BodyParser(&req); err != nil {
		return apperror.BadRequest("invalid request body", err)
	}

	if err := h.validate.Struct(req); err != nil {
		return apperror.BadRequest("invalid request body", err)
	}

	jwtEntity, err := h.jwtService.ParseToken(req.RefreshToken, true)
	if err != nil {
		return apperror.UnAuthorized("invalid token", err)
	}

	cachedToken, err := h.jwtService.GetCachedTokens(ctx, jwtEntity.ID)
	if err != nil {
		return apperror.UnAuthorized("invalid token", err)
	}

	if err := h.jwtService.ValidateToken(*cachedToken, jwtEntity, true); err != nil {
		return apperror.UnAuthorized("invalid token", err)
	}

	tokens, err := h.jwtService.GenerateAndStoreTokenPair(ctx, &model.User{
		Model: gorm.Model{ID: jwtEntity.ID},
		Role:  jwtEntity.Role,
	})
	if err != nil {
		return errors.Wrap(err, "failed to generate token pair")
	}

	return c.Status(fiber.StatusOK).JSON(tokens)
}
