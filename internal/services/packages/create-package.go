package packages

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

func (h *Handler) RegisterCreatePackage(api huma.API, middlewares huma.Middlewares) {
	huma.Register(api, huma.Operation{
		OperationID: "create-package",
		Method:      http.MethodPost,
		Path:        "/api/v1/photographer/packages",
		Summary:     "Create package",
		Description: "Create package",
		Tags:        []string{"packages"},
		Security: []map[string][]string{
			{"bearer": nil},
		},
		Middlewares: middlewares,
	}, h.HandleCreatePackage)
}

func (h *Handler) HandleCreatePackage(ctx context.Context, req *dto.HumaBody[dto.CreatePackageRequest]) (*struct{}, error) {
	userId, err := h.authMiddleware.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user id from context")
	}

	if err := h.validate.Struct(req); err != nil {
		return nil, apperror.BadRequest("invalid request body", err)
	}
	for _, media := range req.Body.Media {
		if err := h.validate.Struct(media); err != nil {
			return nil, apperror.BadRequest("invalid request body", err)
		}
	}

	if req.Body.Price <= 0 {
		return nil, apperror.BadRequest("invalid request body", errors.New("Price must be positive"))
	}

	if err = h.CreatePackage(&req.Body, userId); err != nil {
		return nil, errors.Wrap(err, "failed to create Package")
	}

	return nil, nil
}

func (h *Handler) CreatePackage(req *dto.CreatePackageRequest, userId uint) error {
	if err := h.store.DB.Transaction(func(tx *gorm.DB) error {
		Package := &model.Package{
			PhotographerID: userId,
			Name:           req.Name,
			Description:    req.Description,
			Price:          req.Price,
		}

		if err := tx.Create(&Package).Error; err != nil {
			return errors.Wrap(err, "failed to create Package")
		}

		for _, media := range req.Media {
			if err := tx.Create(&model.Media{
				PictureURL:  media.PictureURL,
				Description: media.Description,
				PackageID:   Package.ID,
			}).Error; err != nil {
				return errors.Wrap(err, "failed to create media")
			}
		}

		return nil
	}); err != nil {
		return errors.Wrap(err, "failed to create Package")
	}

	return nil
}
