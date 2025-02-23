package media

import (
	"context"
	"net/http"

	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/cockroachdb/errors"
	"github.com/danielgtaylor/huma/v2"
	"gorm.io/gorm"
)

func (h *Handler) RegisterCreateMedia(api huma.API, middlewares huma.Middlewares) {
	huma.Register(api, huma.Operation{
		OperationID: "create-media",
		Method:      http.MethodPost,
		Path:        "/api/v1/photographer/media",
		Summary:     "Create media",
		Description: "Create media",
		Tags:        []string{"media"},
		Security: []map[string][]string{
			{"bearer": nil},
		},
		Middlewares: middlewares,
	}, h.HandleCreateMedia)
}

func (h *Handler) HandleCreateMedia(ctx context.Context, req *dto.HumaBody[dto.CreateMediaRequest]) (*struct{}, error) {
	userId, err := h.authMiddleware.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user id from context")
	}

	if err := h.validate.Struct(req); err != nil {
		return nil, huma.Error400BadRequest("invalid request", err)
	}

	if err := h.createMedia(&req.Body, userId); err != nil {
		return nil, errors.Wrap(err, "failed to create media")
	}

	return nil, nil
}

func (h *Handler) createMedia(req *dto.CreateMediaRequest, userId uint) error {
	if err := h.store.DB.Transaction(func(tx *gorm.DB) error {
		var pkg model.Package
		if err := h.store.DB.Preload("Photographer").First(&pkg, "id = ?", req.PackageID).Error; err != nil {
			return errors.Wrap(err, "failed to get package")
		}

		if pkg.Photographer.UserID != userId {
			return apperror.Forbidden("You are not allowed to create media in this package", errors.New("unauthorized"))
		}

		if err := tx.Create(&model.Media{
			PackageID:   req.PackageID,
			PictureURL:  req.PictureURL,
			Description: req.Description,
		}).Error; err != nil {
			return errors.Wrap(err, "failed to create media")
		}

		return nil
	}); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return huma.Error404NotFound("package not found", errors.New("package not found"))
		}
		return errors.Wrap(err, "failed to create media")
	}

	return nil
}
