package media

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func (h *Handler) HandleDeleteMedia(c *fiber.Ctx) error {
	userId, err := h.authMiddleware.GetUserIDFromContext(c.UserContext())
	if err != nil {
		return errors.Wrap(err, "failed to get user id from context")
	}

	req := new(dto.DeleteMediaRequest)
	if err := c.ParamsParser(req); err != nil {
		return apperror.BadRequest("invalid media id", err)
	}

	if err := h.deleteMedia(req.MediaId, userId); err != nil {
		return errors.Wrap(err, "failed to delete media")
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *Handler) deleteMedia(mediaId uint, userId uint) error {
	if err := h.store.DB.Transaction(func(tx *gorm.DB) error {
		var media model.Media
		if err := h.store.DB.Preload("Package.Photographer").First(&media, "id = ?", mediaId).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return apperror.NotFound("Media not found", err)
			}
			return errors.Wrap(err, "Failed to get media")
		}

		if media.Package.Photographer.UserID != userId {
			return apperror.Forbidden("You are not allowed to delete this media", errors.New("unauthorized"))
		}

		if err := h.store.DB.Delete(&media).Error; err != nil {
			return errors.Wrap(err, "failed to delete media")
		}

		return nil
	}); err != nil {
		return errors.Wrap(err, "failed to delete media")
	}

	return nil
}
