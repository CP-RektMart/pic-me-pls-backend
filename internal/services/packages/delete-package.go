package packages

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// @Summary			Delete Package
// @Description		Delete package by photographer
// @Tags			packages
// @Router			/api/v1/photographer/packages/{id} [DELETE]
// @Security		ApiKeyAuth
// @Param        	id 	path     uint    true  "Package ID"
// @Success			204
// @Failure			400	{object}	dto.HttpError
// @Failure			500	{object}	dto.HttpError
func (h *Handler) HandleDeletePackage(c *fiber.Ctx) error {
	userID, err := h.authMiddleware.GetUserIDFromContext(c.UserContext())
	if err != nil {
		return errors.Wrap(err, "Failed to get user id from context")
	}

	req := new(dto.DeletePackageRequest)
	if err := c.ParamsParser(req); err != nil {
		return apperror.BadRequest("Invalid package id", err)
	}

	if err := h.deletePackage(req.ID, userID); err != nil {
		return errors.Wrap(err, "Failed to delete package")
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *Handler) deletePackage(packageID uint, userID uint) error {
	if err := h.store.DB.Transaction(func(tx *gorm.DB) error {
		var pkg model.Package
		if err := h.store.DB.Preload("Media").First(&pkg, "id = ?", packageID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return apperror.NotFound("Package not found", err)
			}
			return errors.Wrap(err, "Failed to get package")
		}

		if pkg.PhotographerID != userID {
			return apperror.Forbidden("You are not allowed to delete this package", errors.New("unauthorized"))
		}

		if err := h.store.DB.Select("Media").Select("Tags").Delete(&pkg).Error; err != nil {
			return errors.Wrap(err, "Failed to delete package")
		}

		return nil
	}); err != nil {
		return errors.Wrap(err, "Failed to delete package")
	}

	return nil
}
