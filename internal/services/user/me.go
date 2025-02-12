package user

import (
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/cockroachdb/errors"

	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/gofiber/fiber/v2"
)

// @Summary			Get me
// @Description		Get me
// @Tags			user
// @Router			/api/v1/me [GET]
// @Security		ApiKeyAuth
// @Success			200	{object}	dto.HttpResponse{result=dto.UserResponse}
// @Failure			400	{object}	dto.HttpResponse
// @Failure			500	{object}	dto.HttpResponse
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

	response := dto.UserResponse{
		ID:                user.ID,
		Name:              user.Name,
		Email:             user.Email,
		ProfilePictureURL: user.ProfilePictureURL,
		Role:              user.Role.String(),
	}

	return c.Status(fiber.StatusOK).JSON(dto.HttpResponse{
		Result: response,
	})
}
