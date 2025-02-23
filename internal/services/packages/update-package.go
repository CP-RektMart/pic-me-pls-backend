package packages

import (
	"context"

	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/cockroachdb/errors"
	"gorm.io/gorm"
)

func (h *Handler) HandleUpdatePackage(ctx context.Context, req *dto.UpdatePackageRequest) (*struct{}, error) {
	userId, err := h.authMiddleware.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user id from context")
	}

	if err := h.validate.Struct(req); err != nil {
		return nil, apperror.BadRequest("invalid request body", err)
	}

	if err := h.updatePackage(&req.Body, req.PackageId, userId); err != nil {
		return nil, errors.Wrap(err, "failed to update gallery")
	}

	return nil, nil
}

func (h *Handler) updatePackage(req *dto.UpdatePackageBody, packageId uint, userId uint) error {
	var pkg model.Package
	if err := h.store.DB.Transaction(func(tx *gorm.DB) error {
		if err := h.store.DB.Preload("Photographer").First(&pkg, "id = ?", packageId).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return apperror.NotFound("Package not found", err)
			}
			return errors.Wrap(err, "Failed to get package")
		}

		if pkg.Photographer.UserID != userId {
			return apperror.Forbidden("You are not allowed to update this package", errors.New("unauthorized"))
		}

		if req.Name != "" {
			pkg.Name = req.Name
		}
		if req.Description != "" {
			pkg.Description = req.Description
		}
		if req.Price != 0 {
			pkg.Price = req.Price
		}

		if err := tx.Save(&pkg).Error; err != nil {
			return errors.Wrap(err, "Failed to update package")
		}

		return nil
	}); err != nil {
		return errors.Wrap(err, "failed to update package")
	}

	return nil
}
