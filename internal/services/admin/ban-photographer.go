package admin

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// @Summary      ban photographer
// @Description  ban photographer by id
// @Tags         admin
// @Router       /api/v1/admin/photographer/{id}/ban [PATCH]
// @Security	 ApiKeyAuth
// @Success			204
// @Failure			400	{object}	dto.HttpError
// @Failure			401	{object}	dto.HttpError
// @Failure			404	{object}	dto.HttpError
// @Failure			500	{object}	dto.HttpError
func (h *Handler) HandleBanPhotographer(c *fiber.Ctx) error {
	var req dto.BanPhotographerRequest
	if err := c.ParamsParser(&req); err != nil {
		return apperror.BadRequest("invalid path params", err)
	}
	if err := h.validate.Struct(req); err != nil {
		return apperror.BadRequest("invalid path params", err)
	}

	if err := h.BanPhotographer(req.ID); err != nil {
		return errors.Wrap(err, "failed to ban photographer")
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *Handler) BanPhotographer(ID uint) error {
	if err := h.store.DB.Transaction(func(tx *gorm.DB) error {
		// DO something
		return nil
	}); err != nil {
		return errors.Wrap(err, "failed to ban photographer")
	}

	return nil
}
