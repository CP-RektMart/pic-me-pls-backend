package packages

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// @Summary			Create Package
// @Description			Create Package by photographer
// @Tags			packages
// @Router			/api/v1/photographer/packages [POST]
// @Security			ApiKeyAuth
// @Param        		RequestBody 	body  dto.CreatePackageRequest  true  "Package details"
// @Success			204
// @Failure			400	{object}	dto.HttpError
// @Failure			500	{object}	dto.HttpError
func (h *Handler) HandleCreatePackage(c *fiber.Ctx) error {
	userID, err := h.authMiddleware.GetUserIDFromContext(c.UserContext())
	if err != nil {
		return errors.Wrap(err, "failed to get user id from context")
	}

	var req dto.CreatePackageRequest
	if err := c.BodyParser(&req); err != nil {
		return apperror.BadRequest("invalid request body", err)
	}

	if err := h.validate.Struct(req); err != nil {
		return apperror.BadRequest("invalid request body", err)
	}

	for _, media := range req.Media {
		if err := h.validate.Struct(media); err != nil {
			return apperror.BadRequest("invalid request body", err)
		}
	}

	if err = h.createPackage(&req, userID); err != nil {
		return errors.Wrap(err, "failed to create Package")
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *Handler) createPackage(req *dto.CreatePackageRequest, photographerID uint) error {

	// Validate the request
	if err := ValidateCreatePackageRequest(req, photographerID); err != nil {
		return errors.Wrap(err, "Invalid request")
	}

	newPackage := &model.Package{
		Name:           req.Name,
		Description:    req.Description,
		Price:          req.Price,
		PhotographerID: photographerID,
		CategoryID:     req.CategoryID,
		Media:          dto.ToPackageMediaModels(req.Media),
	}

	var photographer model.Photographer
	if err := h.store.DB.First(&photographer, photographerID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return apperror.NotFound("Photographer not found", err)
		}
		return errors.Wrap(err, "failed to get photographer")
	}

	if photographer.IsBanned {
		return apperror.Forbidden("You are banned from creating packages", errors.New("banned photographer"))
	}

	if err := h.store.DB.Create(&newPackage).Error; err != nil {
		return errors.Wrap(err, "failed to save Package to DB")
	}

	return nil
}

func ValidateCreatePackageRequest(req *dto.CreatePackageRequest, photographerID uint) error {
	if req.Name == "" {
		return apperror.BadRequest("Package name is required", nil)
	}
	if req.Description == "" {
		return apperror.BadRequest("Package description is required", nil)
	}
	if req.Price <= 0 {
		return apperror.BadRequest("Package price must be greater than 0", nil)
	}
	if photographerID <= 0 {
		return apperror.BadRequest("Photographer ID must be greater than 0", nil)
	}
	if len(req.Media) == 0 {
		return apperror.BadRequest("At least one media file is required", nil)
	}

	return nil
}
