package admin

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// @Summary      Verify photographer
// @Description  Verify a photographer by their ID
// @Tags         admin
// @Router       /api/v1/admin/photographers/{photographerID}/verify [PATCH]
// @Security	 ApiKeyAuth
// @Param        id          path     int  true   "{Photographer Id}"
// @Success      204            "No Content"
// @Failure      400            {object}  dto.HttpError
// @Failure      404            {object}  dto.HttpError
// @Failure      500            {object}  dto.HttpError
func (h *Handler) HandleVerifyPhotographer(c *fiber.Ctx) error {
	var req dto.VerifyPhotographerRequest

	if err := c.ParamsParser(&req); err != nil {
		return apperror.BadRequest("invalid path params", err)
	}

	if err := h.validate.Struct(req); err != nil {
		return apperror.BadRequest("invalid path params", err)
	}

	if err := h.verifyPhotographer(req.PhotographerID); err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *Handler) verifyPhotographer(photographerId uint) error {
    result := h.store.DB.Model(&model.Photographer{}).
        Where("user_id = ?", photographerId).
        Update("is_verified", true)

    if result.Error != nil {
        return errors.Wrap(result.Error, "failed to verify photographer")
    }

    if result.RowsAffected == 0 {
        return apperror.NotFound("Photographer not found", gorm.ErrRecordNotFound)
    }

    return nil
}