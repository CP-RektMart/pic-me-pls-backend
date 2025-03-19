package media

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// @Summary			Delete Media
// @Description		Delete media
// @Tags			media
// @Router			/api/v1/photographer/media/{mediaId} [DELETE]
// @Security		ApiKeyAuth
// @Param 	 		mediaId			path      string    	true  "media id"
// @Param        	RequestBody 	body  dto.DeleteMediaRequest  true  "Media details"
// @Success			204
// @Failure			400	{object}	dto.HttpError
// @Failure			500	{object}	dto.HttpError
func (h *Handler) HandleDeleteMedia(c *fiber.Ctx) error {
	userID, err := h.authMiddleware.GetUserIDFromContext(c.UserContext())
	if err != nil {
		return errors.Wrap(err, "failed to get user id from context")
	}

	req := new(dto.DeleteMediaRequest)
	if err := c.ParamsParser(req); err != nil {
		return apperror.BadRequest("invalid media id", err)
	}

	if err := h.deleteMedia(req.MediaID, userID); err != nil {
		return errors.Wrap(err, "failed to delete media")
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *Handler) deleteMedia(mediaID uint, userID uint) error {
	if err := h.store.DB.Transaction(func(tx *gorm.DB) error {
		var media model.Media
		if err := h.store.DB.Preload("Package").First(&media, "id = ?", mediaID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return apperror.NotFound("Media not found", err)
			}
			return errors.Wrap(err, "Failed to get media")
		}

		if media.Package.PhotographerID != userID {
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
