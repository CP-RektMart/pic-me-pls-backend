package admin

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// @Summary			admin delete package
// @Tags			admin
// @Router			/api/v1/admin/packages/{packageID} [DELETE]
// @Security		ApiKeyAuth
// @Param			packageID		path		int	true	"package id"
// @Success			204
// @Failure			400	{object}	dto.HttpError
// @Failure			401	{object}	dto.HttpError
// @Failure			403	{object}	dto.HttpError
// @Failure			404	{object}	dto.HttpError
// @Failure			500	{object}	dto.HttpError
func (h *Handler) HandleDeletePackageByID(c *fiber.Ctx) error {
	var req dto.AdminDeletePackageByID
	if err := c.ParamsParser(&req); err != nil {
		return apperror.BadRequest("invalid request", err)
	}
	if err := h.validate.Struct(req); err != nil {
		return apperror.BadRequest("invalid request", err)
	}

	if err := h.deletePackageByID(req.PackageID); err != nil {
		return errors.Wrap(err, "failed delete package")
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *Handler) deletePackageByID(ID uint) error {
	if err := h.store.DB.Delete(&model.Package{Model: gorm.Model{ID: ID}}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return apperror.NotFound("package not found", err)
		}
		return err
	}
	return nil
}
