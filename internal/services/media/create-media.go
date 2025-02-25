package media

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// @Summary			Create Media
// @Description		Create media by photographer
// @Tags			media
// @Router			/api/v1/photographer/media [POST]
// @Security		ApiKeyAuth
// @Param        	RequestBody 	body  dto.CreateMediaRequest  true  "Media details"
// @Success			201
// @Failure			400	{object}	dto.HttpError
// @Failure			500	{object}	dto.HttpError
func (h *Handler) HandleCreateMedia(c *fiber.Ctx) error {
	userId, err := h.authMiddleware.GetUserIDFromContext(c.UserContext())
	if err != nil {
		return errors.Wrap(err, "failed to get user id from context")
	}

	req := new(dto.CreateMediaRequest)
	if err := c.BodyParser(req); err != nil {
		return apperror.BadRequest("invalid request body", err)
	}

	if err := h.validate.Struct(req); err != nil {
		return apperror.BadRequest("invalid request body", err)
	}

	if err := h.createMedia(req, userId); err != nil {
		return errors.Wrap(err, "failed to create media")
	}

	return c.SendStatus(fiber.StatusCreated)
}

func (h *Handler) createMedia(req *dto.CreateMediaRequest, userId uint) error {
	if err := h.store.DB.Transaction(func(tx *gorm.DB) error {
		var pkg model.Package
		if err := h.store.DB.Preload("Photographer").First(&pkg, "id = ?", req.PackageID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return apperror.NotFound("Package not found", err)
			}
			return errors.Wrap(err, "failed to get package")
		}

		if pkg.Photographer.UserID != userId {
			return apperror.Forbidden("You are not allowed to create media in this package", errors.New("unauthorized"))
		}

		if err := tx.Create(&model.Media{
			PackageID:   req.PackageID,
			PictureURL:  req.PictureURL,
			Description: req.Description,
		}).Error; err != nil {
			return errors.Wrap(err, "failed to create media")
		}

		return nil
	}); err != nil {
		return errors.Wrap(err, "failed to create media")
	}

	return nil
}
