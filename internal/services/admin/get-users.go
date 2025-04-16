package admin

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// @Summary      get all users
// @Description  get all users
// @Tags         admin
// @Router       /api/v1/admin/user [GET]
// @Security			ApiKeyAuth
// @Success      200    {object}  dto.HttpListResponse[dto.PublicUserResponse]
// @Failure      400    {object}  dto.HttpError
// @Failure      404    {object}  dto.HttpError
// @Failure      500    {object}  dto.HttpError
func (h *Handler) HandleGetAllUsers(c *fiber.Ctx) error {
	var users []model.User

	if err := h.store.DB.Find(&users).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return apperror.NotFound("users not found", err)
		}

		return errors.Wrap(err, "failed fetch users")
	}

	return c.Status(fiber.StatusOK).JSON(dto.HttpListResponse[dto.PublicUserResponse]{
		Result: dto.ToPublicUserResponses(users),
	})
}
