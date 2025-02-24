package packages

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// @Summary			Update package
// @Description		Update
// @Tags			packages
// @Router			/api/v1/packages/{packageId} [PATCH]
// @Security		ApiKeyAuth
// @Param        	RequestBody 	body  dto.UpdatePackageRequest  true  "Package details"
// @Success			204
// @Failure			400	{object}	dto.HttpError
// @Failure			500	{object}	dto.HttpError
func (h *Handler) HandleUpdatePackage(c *fiber.Ctx) error {
	userId, err := h.authMiddleware.GetUserIDFromContext(c.UserContext())
	if err != nil {
		return errors.Wrap(err, "failed to get user id from context")
	}

	req := new(dto.UpdatePackageRequest)
	if err := c.ParamsParser(req); err != nil {
		return apperror.BadRequest("invalid package id", err)
	}

	if err := c.BodyParser(req); err != nil {
		return apperror.BadRequest("invalid request body", err)
	}

	if err := h.validate.Struct(req); err != nil {
		return apperror.BadRequest("invalid request body", err)
	}

	if err := h.updatePackage(req, req.PackageId, userId); err != nil {
		return errors.Wrap(err, "failed to update package")
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *Handler) updatePackage(req *dto.UpdatePackageRequest, packageId uint, userId uint) error {
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
