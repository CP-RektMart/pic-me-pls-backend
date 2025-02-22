package gallery

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// @Summary			Update gallery
// @Description		Update
// @Tags			gallery
// @Router			/api/v1/gallery [PATCH]
// @Security		ApiKeyAuth
// @Param        	RequestBody 	body  dto.UpdateGalleryRequest  true  "Gallery details"
// @Success			200 {object}	dto.HttpResponse[dto.UpdateGalleryResponse]
// @Failure			400	{object}	dto.HttpError
// @Failure			500	{object}	dto.HttpError
func (h *Handler) HandleUpdateGallery(c *fiber.Ctx) error {
	userId, err := h.authMiddleware.GetUserIDFromContext(c.UserContext())
	if err != nil {
		return errors.Wrap(err, "failed to get user id from context")
	}

	req := new(dto.UpdateGalleryRequest)
	if err := c.ParamsParser(req); err != nil {
		return apperror.BadRequest("invalid gallery id", err)
	}

	if err := c.BodyParser(req); err != nil {
		return apperror.BadRequest("invalid request body", err)
	}

	if err := h.validate.Struct(req); err != nil {
		return apperror.BadRequest("invalid request body", err)
	}

	gallery, err := h.updateGallery(req, req.GalleryId, userId)
	if err != nil {
		return errors.Wrap(err, "failed to update gallery")
	}

	return c.Status(fiber.StatusOK).JSON(dto.HttpResponse[dto.UpdateGalleryResponse]{
		Result: dto.UpdateGalleryResponse{
			Name:        gallery.Name,
			Description: gallery.Description,
			Price:       gallery.Price,
		},
	})
}

func (h *Handler) updateGallery(req *dto.UpdateGalleryRequest, galleryId uint, userId uint) (*model.Gallery, error) {
	var gallery model.Gallery
	if err := h.store.DB.Transaction(func(tx *gorm.DB) error {
		if err := h.store.DB.Preload("Photographer").First(&gallery, "id = ?", galleryId).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return apperror.NotFound("Gallery not found", errors.New("Gallery not found"))
			}
			return errors.Wrap(err, "Failed to get gallery")
		}

		if gallery.Photographer.UserID != userId {
			return apperror.Forbidden("You are not allowed to update this gallery", errors.New("unauthorized"))
		}

		if req.Name != "" {
			gallery.Name = req.Name
		}
		if req.Description != "" {
			gallery.Description = req.Description
		}
		if req.Price != 0 {
			gallery.Price = req.Price
		}

		if err := tx.Save(&gallery).Error; err != nil {
			return errors.Wrap(err, "Failed to update gallery")
		}

		return nil
	}); err != nil {
		return nil, errors.Wrap(err, "failed to update gallery")
	}

	return &gallery, nil
}
