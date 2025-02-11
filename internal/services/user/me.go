package user

import (
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/cockroachdb/errors"

	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/gofiber/fiber/v2"
)

// handlerGetMe godoc
// @summary Get user profile
// @description Retrieves the authenticated user's profile.
// @tags user
// @security Bearer
// @id get-me
// @accept json
// @produce json
// @success 200 {object} dto.BaseUserDTO "OK"
// @failure 400 {object} dto.HttpResponse "Bad Request"
// @failure 500 {object} dto.HttpResponse "Internal Server Error"
// @Router /api/v1/me [GET]
func (h *Handler) HandleGetMe(c *fiber.Ctx) error {
	userId, err := h.authMiddleware.GetUserIDFromContext(c.UserContext())
	if err != nil {
		return errors.Wrap(err, "failed to get user id from context")
	}

	var user model.User
	result := h.store.DB.First(&user, userId)
	if result.Error != nil {
		return apperror.Internal("failed to get user", nil)
	}

	userDTO := dto.BaseUserDTO{
		ID:                user.ID,
		Name:              user.Name,
		Email:             user.Email,
		ProfilePictureURL: user.ProfilePictureURL,
		Role:              user.Role.String(),
	}

	return c.Status(fiber.StatusOK).JSON(dto.HttpResponse{Result: userDTO})
}
