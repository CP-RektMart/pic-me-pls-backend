package photographers

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// @Summary      get photographer by id
// @Description  get photographer by id
// @Tags         photographers
// @Router       /api/v1/photographers/{id} [GET]
// @Param 	 id	path      uint    	true  "photographer id"
// @Success      200    {object}  dto.HttpResponse[dto.PhotographerResponse]
// @Failure      400    {object}  dto.HttpError
// @Failure      404    {object}  dto.HttpError
// @Failure      500    {object}  dto.HttpError
func (h *Handler) HandlerGetPhotographerByID(c *fiber.Ctx) error {
	var req dto.GetPhotographerByIDRequest
	if err := c.ParamsParser(&req); err != nil {
		return apperror.BadRequest("invalid path params", err)
	}
	if err := h.validate.Struct(req); err != nil {
		return apperror.BadRequest("invalid path params", err)
	}

	photographer, err := h.getPhotographerByID(req.ID)
	if err != nil {
		return errors.Wrap(err, "failed get photographer by id")
	}

	return c.Status(fiber.StatusOK).JSON(dto.HttpResponse[dto.PhotographerResponse]{
		Result: dto.ToPhotographerResponse(photographer),
	})
}

func (h *Handler) getPhotographerByID(ID uint) (model.Photographer, error) {
	var photographer model.Photographer
	if err := h.store.DB.Preload("User").First(&photographer, ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return photographer, apperror.NotFound("Photographer not found", err)
		}
		return photographer, errors.Wrap(err, "failed fetch photographer")
	}
	return photographer, nil
}
