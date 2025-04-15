package admin

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// @Summary			get photographers detail
// @Tags			admin
// @Router			/api/v1/admin/photographers/{id} [GET]
// @Security		ApiKeyAuth
// @Param			photographerID		path		int	true	"photographer id"
// @Success			200	{object}	dto.HttpResponse[dto.ListPackageResponse]
// @Failure			400	{object}	dto.HttpError
// @Failure			401	{object}	dto.HttpError
// @Failure			403	{object}	dto.HttpError
// @Failure			404	{object}	dto.HttpError
// @Failure			500	{object}	dto.HttpError
func (h *Handler) HandleGetPhotographerByID(c *fiber.Ctx) error {
	var req dto.AdminGetPhotographerByID
	if err := c.ParamsParser(&req); err != nil {
		return apperror.BadRequest("invalid request", err)
	}
	if err := h.validate.Struct(req); err != nil {
		return apperror.BadRequest("invalid request", err)
	}

	photographer, err := h.getPhotographerByID(req.PhotographerID)
	if err != nil {
		return err
	}

	result := dto.ToDetailPhotographerResponse(photographer)

	return c.Status(fiber.StatusOK).JSON(dto.Success(result))
}

func (h *Handler) getPhotographerByID(ID uint) (model.Photographer, error) {
	var p model.Photographer
	if err := h.store.DB.Preload("Packages.Category").Preload("Packages.Tags").Preload("User").First(&p, ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Photographer{}, apperror.NotFound("photographer not found", err)
		}
		return model.Photographer{}, err
	}

	return p, nil
}
