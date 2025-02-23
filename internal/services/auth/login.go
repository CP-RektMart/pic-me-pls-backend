package auth

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// @Summary			Login
// @Description		Login
// @Tags			auth
// @Router			/api/v1/auth/login [POST]
// @Param 			RequestBody 	body 	dto.LoginRequest 	true 	"request request"
// @Success			200	{object}	dto.HttpResponse[dto.LoginResponse]
// @Failure			400	{object}	dto.HttpError
// @Failure			500	{object}	dto.HttpError
func (h *Handler) HandleLogin(c *fiber.Ctx) error {
	ctx := c.UserContext()

	req := new(dto.LoginRequest)
	if err := c.BodyParser(req); err != nil {
		return apperror.BadRequest("invalid request body", err)
	}

	if err := h.validate.Struct(req); err != nil {
		return apperror.BadRequest("invalid request body", err)
	}

	OAuthUser, err := h.validateIDToken(ctx, req.IDToken)
	if err != nil {
		return errors.Wrap(err, "failed to validate id token")
	}

	var user *model.User
	var token *model.Token

	if err := h.store.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("email = ?", OAuthUser.Email).First(&user).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return apperror.NotFound("account not found", err)
			}
			return errors.Wrap(err, "failed getting user")
		}

		token, err = h.jwtService.GenerateAndStoreTokenPair(ctx, user)
		if err != nil {
			return errors.Wrap(err, "failed to generate token pair")
		}

		return nil
	}); err != nil {
		return errors.Wrap(err, "failed to get user and token")
	}

	result := dto.LoginResponse{
		TokenResponse: dto.TokenResponse{
			AccessToken:  token.AccessToken,
			RefreshToken: token.RefreshToken,
			Exp:          token.Exp,
		},
		User: dto.ToUserResponse(*user),
	}

	return c.Status(fiber.StatusOK).JSON(dto.HttpResponse[dto.LoginResponse]{
		Result: result,
	})
}
