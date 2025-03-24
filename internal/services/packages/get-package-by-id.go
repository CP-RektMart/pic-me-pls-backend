package packages

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// @Summary      get package by id
// @Description  Show package detail
// @Tags         packages
// @Router       /api/v1/packages/{id} [GET]
// @Param 	 id	path      string    	true  "package id"
// @Success      200    {object}  dto.HttpResponse[dto.PackageResponse]
// @Failure      400    {object}  dto.HttpError
// @Failure      500    {object}  dto.HttpError
func (h *Handler) HandleGetPackageByID(c *fiber.Ctx) error {
	var req dto.GetPackageByIDRequest
	if err := c.ParamsParser(&req); err != nil {
		return apperror.BadRequest("invalid path params", err)
	}
	if err := h.validate.Struct(req); err != nil {
		return apperror.BadRequest("invalid path params", err)
	}	

	pkg, err := h.getPackageByID(req.ID)
	if err != nil {
		return errors.Wrap(err, "failed get package by id")
	}

	return c.Status(fiber.StatusOK).JSON(dto.HttpResponse[dto.PackageResponse]{
		Result: dto.ToPackageResponse(pkg),
	})
}

func (h *Handler) getPackageByID(ID uint) (model.Package, error) {
	var pkg model.Package
	if err := h.store.DB.
		Preload("Category").
		Preload("Photographer.User").
		Preload("Tags").
		Preload("Reviews.Customer").
		Preload("Media").First(&pkg, ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return model.Package{}, apperror.NotFound("package not found", err)
		}
		return model.Package{}, errors.Wrap(err, "Failed fetch package")
	}

	return pkg, nil
}
