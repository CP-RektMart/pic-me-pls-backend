package admin

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// @Summary      get user by id
// @Description  get user by id for admin
// @Tags         admin
// @Router       /api/v1/admin/user/{id} [GET]
// @Security			ApiKeyAuth
// @Success      200    {object}  dto.HttpResponse[dto.PublicUserResponse]
// @Failure      400    {object}  dto.HttpError
// @Failure      404    {object}  dto.HttpError
// @Failure      500    {object}  dto.HttpError
func (h *Handler) HandleGetUserByID(c *fiber.Ctx) error {
	var req dto.GetPhotographerByIDRequest
	if err := c.ParamsParser(&req); err != nil {
		return apperror.BadRequest("invalid path params", err)
	}
	if err := h.validate.Struct(req); err != nil {
		return apperror.BadRequest("invalid path params", err)
	}

	user, err := h.getUserByID(req.ID)
	if err != nil {
		return errors.Wrap(err, "failed get user by id")
	}

	return c.Status(fiber.StatusOK).JSON(dto.HttpResponse[dto.PublicUserResponse]{
		Result: dto.ToPublicUserResponse(user),
	})
}

func (h *Handler) getUserByID(ID uint) (model.User, error) {
	var user model.User
	if err := h.store.DB.First(&user, ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return user, apperror.NotFound("User not found", err)
		}
		return user, errors.Wrap(err, "failed fetch user")
	}
	return user, nil
}
