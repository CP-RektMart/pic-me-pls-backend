package packages

import (
	"context"

	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/cockroachdb/errors"
	"gorm.io/gorm"
)

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
