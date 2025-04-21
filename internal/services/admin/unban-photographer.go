package admin

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// @Summary      unban photographer
// @Description  unban photographer by id
// @Tags         admin
// @Router       /api/v1/admin/photographer/{id}/unban [PATCH]
// @Security	 ApiKeyAuth
// @Param        id          path     int  true   "Photographer ID"
// @Success			204
// @Failure			400	{object}	dto.HttpError
// @Failure			401	{object}	dto.HttpError
// @Failure			404	{object}	dto.HttpError
// @Failure			500	{object}	dto.HttpError
func (h *Handler) HandleUnbanPhotographer(c *fiber.Ctx) error {
	var req dto.UnbanPhotographerRequest
	if err := c.ParamsParser(&req); err != nil {
		return apperror.BadRequest("invalid path params", err)
	}

	if err := h.validate.Struct(req); err != nil {
		return apperror.BadRequest("invalid path params", err)
	}

	if err := h.UnbanPhotographer(req.ID); err != nil {
		return errors.Wrap(err, "failed to unban photographer")
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *Handler) UnbanPhotographer(ID uint) error {
	if err := h.store.DB.Transaction(func(tx *gorm.DB) error {
		var photographer model.Photographer
		if err := tx.First(&photographer, ID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return apperror.NotFound("photographer not found", err)
			}
			return errors.Wrap(err, "failed to fetch photographer")
		}

		if !photographer.IsBanned {
			return apperror.BadRequest("photographer is not currently banned", nil)
		}

		photographer.IsBanned = false

		if err := tx.Save(&photographer).Error; err != nil {
			return errors.Wrap(err, "failed to unban photographer")
		}

		return nil
	}); err != nil {
		return errors.Wrap(err, "failed to unban photographer")
	}

	return nil
}
