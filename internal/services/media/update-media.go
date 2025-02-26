package media

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// @Summary			Update Media
// @Description		Update media
// @Tags			media
// @Router			/api/v1/photographer/media/{mediaId} [PATCH]
// @Security		ApiKeyAuth
// @Param        	RequestBody 	body  dto.UpdateMediaRequest  true  "Media details"
// @Success			204
// @Failure			400	{object}	dto.HttpError
// @Failure			500	{object}	dto.HttpError
func (h *Handler) HandleUpdateMedia(c *fiber.Ctx) error {
	userID, err := h.authMiddleware.GetUserIDFromContext(c.UserContext())
	if err != nil {
		return errors.Wrap(err, "failed to get user id from context")
	}

	req := new(dto.UpdateMediaRequest)
	if err = c.ParamsParser(req); err != nil {
		return apperror.BadRequest("invalid media id", err)
	}

	if err = c.BodyParser(req); err != nil {
		return apperror.BadRequest("invalid request body", err)
	}

	if err := h.validate.Struct(req); err != nil {
		return apperror.BadRequest("invalid request body", err)
	}

	if err = h.updateMedia(req, req.MediaID, userID); err != nil {
		return errors.Wrap(err, "failed to update media")
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *Handler) updateMedia(req *dto.UpdateMediaRequest, mediaID uint, userID uint) error {
	var media model.Media
	if err := h.store.DB.Transaction(func(tx *gorm.DB) error {
		if err := h.store.DB.Preload("Package.Photographer").First(&media, "id = ?", mediaID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return apperror.NotFound("Media not found", err)
			}
			return errors.Wrap(err, "Failed to get media")
		}

		if media.Package.Photographer.UserID != userID {
			return apperror.Forbidden("You are not allowed to update this media", errors.New("unauthorized"))
		}

		if req.PictureURL != "" {
			media.PictureURL = req.PictureURL
		}

		if req.Description != "" {
			media.Description = req.Description
		}

		if err := tx.Save(&media).Error; err != nil {
			return errors.Wrap(err, "Failed to update media")
		}

		return nil
	}); err != nil {
		return errors.Wrap(err, "failed to update media")
	}

	return nil

}
