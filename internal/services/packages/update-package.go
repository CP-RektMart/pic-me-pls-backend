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
// @Description			Update package
// @Tags			packages
// @Router			/api/v1/photographer/packages/{id} [PATCH]
// @Security			ApiKeyAuth
// @Param        		RequestBody 	body  dto.UpdatePackageRequest  true  "Package details"
// @Param			id	path uint true "package id"
// @Success			204
// @Failure			400	{object}	dto.HttpError
// @Failure			500	{object}	dto.HttpError
func (h *Handler) HandleUpdatePackage(c *fiber.Ctx) error {
	userID, err := h.authMiddleware.GetUserIDFromContext(c.UserContext())
	if err != nil {
		return errors.Wrap(err, "failed to get user id from context")
	}
	var req dto.UpdatePackageRequest
	if err := c.ParamsParser(&req); err != nil {
		return apperror.BadRequest("invalid path params", err)
	}

	if err := c.BodyParser(&req); err != nil {
		return apperror.BadRequest("invalid request body", err)
	}

	if err := h.validate.Struct(req); err != nil {
		return apperror.BadRequest("invalid request body", err)
	}

	if err := h.updatePackage(&req, req.ID, userID); err != nil {
		return errors.Wrap(err, "failed to update package")
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *Handler) updatePackage(req *dto.UpdatePackageRequest, packageID uint, photographerID uint) error {
	var pkg model.Package

	if err := h.store.DB.Transaction(func(tx *gorm.DB) error {
		if err := h.store.DB.First(&pkg, packageID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return apperror.NotFound("Package not found", err)
			}
			return errors.Wrap(err, "Failed to get package")
		}

		if pkg.PhotographerID != photographerID {
			return apperror.Forbidden("You are not allowed to update this package", errors.New("FORBIDDEN"))
		}

		if err := h.store.DB.Model(&pkg).Updates(&req).Error; err != nil {
			return errors.Wrap(err, "Failed update package")
		}

		return nil
	}); err != nil {
		return errors.Wrap(err, "failed to update package")
	}

	return nil
}
