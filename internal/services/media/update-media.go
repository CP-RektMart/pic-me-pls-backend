package media

import (
	"context"

	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/cockroachdb/errors"
	"github.com/danielgtaylor/huma/v2"
	"gorm.io/gorm"
)

var (
	ErrorMediaNotAllowed = errors.New("MEDIA_NOT_ALLOWED")
)

func (h *Handler) HandleUpdateMedia(ctx context.Context, req *dto.UpdateMediaRequest) (*dto.HumaHttpResponse[dto.CitizenCardResponse], error) {
	userId, err := h.authMiddleware.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user id from context")
	}

	if err := h.validate.Struct(req); err != nil {
		return nil, apperror.BadRequest("invalid request body", err)
	}

	if err = h.updateMedia(&req.Body, req.MediaId, userId); err != nil {
		return nil, errors.Wrap(err, "failed to update media")
	}

	return nil, nil
}

func (h *Handler) updateMedia(req *dto.UpdateMediaBody, mediaId uint, userId uint) error {
	var media model.Media
	if err := h.store.DB.Transaction(func(tx *gorm.DB) error {
		if err := h.store.DB.Preload("Package.Photographer").First(&media, "id = ?", mediaId).Error; err != nil {
			return errors.Wrap(err, "Failed to get media")
		}

		if media.Package.Photographer.UserID != userId {
			return errors.WithStack(ErrorMediaNotAllowed)
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
		if errors.Is(err, ErrorMediaNotAllowed) {
			return huma.Error403Forbidden("You are not allowed to update this media", errors.New("unauthorized"))
		}
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return huma.Error404NotFound("Media not found", err)
		}
		return errors.Wrap(err, "failed to update media")
	}

	return nil

}
