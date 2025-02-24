package media

import (
	"context"
	"net/http"

	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/cockroachdb/errors"
	"github.com/danielgtaylor/huma/v2"
	"gorm.io/gorm"
)

func (h *Handler) RegisterDeleteMedia(api huma.API, middlewares ...func(ctx huma.Context, next func(huma.Context))) {
	huma.Register(api, huma.Operation{
		OperationID: "delete-media",
		Method:      http.MethodDelete,
		Path:        "/api/v1/photographer/media/{mediaId}",
		Summary:     "Delete media",
		Description: "Delete media",
		Tags:        []string{"media"},
		Security: []map[string][]string{
			{"bearer": nil},
		},
		Middlewares: middlewares,
	}, h.HandleDeleteMedia)
}

func (h *Handler) HandleDeleteMedia(ctx context.Context, req *dto.DeleteMediaRequest) (*dto.HumaHttpResponse[dto.CitizenCardResponse], error) {
	userId, err := h.authMiddleware.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user id from context")
	}

	if err := h.deleteMedia(req.MediaId, userId); err != nil {
		return nil, errors.Wrap(err, "failed to delete media")
	}

	return nil, nil
}

func (h *Handler) deleteMedia(mediaId uint, userId uint) error {
	if err := h.store.DB.Transaction(func(tx *gorm.DB) error {
		var media model.Media
		if err := h.store.DB.Preload("Package.Photographer").First(&media, "id = ?", mediaId).Error; err != nil {
			return errors.Wrap(err, "Failed to get media")
		}

		if media.Package.Photographer.UserID != userId {
			return errors.WithStack(ErrorMediaNotAllowed)
		}

		if err := h.store.DB.Delete(&media).Error; err != nil {
			return errors.Wrap(err, "failed to delete media")
		}

		return nil
	}); err != nil {
		if errors.Is(err, ErrorMediaNotAllowed) {
			return huma.Error403Forbidden("You are not allowed to delete this media", errors.New("unauthorized"))
		}
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return huma.Error403Forbidden("Media not found", err)
		}
		return errors.Wrap(err, "failed to delete media")
	}

	return nil
}
