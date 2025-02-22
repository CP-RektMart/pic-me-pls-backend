package packages

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// @Summary			Create Package
// @Description		Create Package by photographer
// @Tags			package
// @Router			/api/v1/Package [POST]
// @Security		ApiKeyAuth
// @Param        	RequestBody 	body  dto.CreatePackageRequest  true  "Package details"
// @Success			201
// @Failure			400	{object}	dto.HttpError
// @Failure			500	{object}	dto.HttpError
func (h *Handler) HandleCreatePackage(c *fiber.Ctx) error {
	userId, err := h.authMiddleware.GetUserIDFromContext(c.UserContext())
	if err != nil {
		return errors.Wrap(err, "failed to get user id from context")
	}

	req := new(dto.CreatePackageRequest)
	if err := c.BodyParser(req); err != nil {
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

	if req.Price <= 0 {
		return apperror.BadRequest("invalid request body", errors.New("Price must be positive"))
	}

	if err = h.CreatePackage(req, userId); err != nil {
		return errors.Wrap(err, "failed to create Package")
	}

	return c.SendStatus(fiber.StatusCreated)
}

func (h *Handler) CreatePackage(req *dto.CreatePackageRequest, userId uint) error {
	if err := h.store.DB.Transaction(func(tx *gorm.DB) error {
		Package := &model.Package{
			PhotographerID: userId,
			Name:           req.Name,
			Description:    req.Description,
			Price:          req.Price,
		}

		if err := tx.Create(&Package).Error; err != nil {
			return errors.Wrap(err, "failed to create Package")
		}

		for _, media := range req.Media {
			if err := tx.Create(&model.Media{
				PictureURL:  media.PictureURL,
				Description: media.Description,
				PackageID:   Package.ID,
			}).Error; err != nil {
				return errors.Wrap(err, "failed to create media")
			}
		}

		return nil
	}); err != nil {
		return errors.Wrap(err, "failed to create Package")
	}

	return nil
}
