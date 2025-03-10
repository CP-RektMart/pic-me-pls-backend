package packages

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
)

// @Summary			Create Package
// @Description			Create Package by photographer
// @Tags			packages
// @Router			/api/v1/photographer/packages [POST]
// @Security			ApiKeyAuth
// @Param        		RequestBody 	body  dto.CreatePackageRequest  true  "Package details"
// @Success			201
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

	return c.SendStatus(fiber.StatusCreated)
}

func (h *Handler) createPackage(req *dto.CreatePackageRequest, userID uint) error {
	var photographer model.Photographer
	if err := h.store.DB.Where("user_id = ?", userID).First(&photographer).Error; err != nil {
		return errors.Wrap(err, "Failed fetch photographer")
	}

	newPackage := &model.Package{
		Name:           req.Name,
		Description:    req.Description,
		Price:          req.Price,
		PhotographerID: photographer.ID,
		CategoryID:     req.CategoryID,
		Media:          dto.ToPackageMediaModels(req.Media),
	}
	if err := h.store.DB.Create(&newPackage).Error; err != nil {
		return errors.Wrap(err, "failed to save Package to DB")
	}

	return nil
}
