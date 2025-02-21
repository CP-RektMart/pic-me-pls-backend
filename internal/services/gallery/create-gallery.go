package gallery

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// @Summary			Create gallery
// @Description		Create gallery by photographer
// @Tags			gallery
// @Router			/api/v1/gallery [POST]
// @Security		ApiKeyAuth
// @Param        	RequestBody 	body  dto.CreateGalleryRequest  true  "Gallery details"
// @Success			201
// @Failure			400	{object}	dto.HttpError
// @Failure			500	{object}	dto.HttpError
func (h *Handler) HandleCreateGallery(c *fiber.Ctx) error {
	userId, err := h.authMiddleware.GetUserIDFromContext(c.UserContext())
	if err != nil {
		return errors.Wrap(err, "failed to get user id from context")
	}

	req := new(dto.CreateGalleryRequest)
	if err := c.BodyParser(req); err != nil {
		return apperror.BadRequest("invalid request body", err)
	}

	if err := h.validate.Struct(req); err != nil {
		return apperror.BadRequest("invalid request body", err)
	}

	if req.Price <= 0 {
		return apperror.BadRequest("invalid request body", errors.New("Price must be positive"))
	}

	if err = h.CreateGallery(req, userId); err != nil {
		return errors.Wrap(err, "failed to create gallery")
	}

	return c.SendStatus(fiber.StatusCreated)
}

func (h *Handler) CreateGallery(req *dto.CreateGalleryRequest, userId uint) error {
	if err := h.store.DB.Transaction(func(tx *gorm.DB) error {
		gallery := &model.Gallery{
			PhotographerID: userId,
			Name:           req.Name,
			Description:    req.Description,
			Price:          req.Price,
		}

		if err := tx.Create(&gallery).Error; err != nil {
			return errors.Wrap(err, "failed to create gallery")
		}

		for _, media := range req.Media {
			if err := tx.Create(&model.Media{
				PictureURL:  media.PictureURL,
				Description: media.Description,
				GalleryID:   gallery.ID,
			}).Error; err != nil {
				return errors.Wrap(err, "failed to create media")
			}
		}

		return nil
	}); err != nil {
		return errors.Wrap(err, "failed to create gallery")
	}

	return nil
}
