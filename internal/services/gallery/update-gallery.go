package gallery

import (
	"strconv"

	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func (h *Handler) HandleUpdateGallery(c *fiber.Ctx) error {
	userId, err := h.authMiddleware.GetUserIDFromContext(c.UserContext())
	if err != nil {
		return errors.Wrap(err, "failed to get user id from context")
	}

	galleryId, err := strconv.Atoi(c.Params("galleryId"))
	if err != nil {
		return apperror.BadRequest("invalid gallery id", errors.Errorf("gallery id is required"))
	}

	req := new(dto.UpdateGalleryRequest)
	if err := c.BodyParser(req); err != nil {
		return apperror.BadRequest("invalid request body", err)
	}

	if req.Price < 0 {
		return apperror.BadRequest("invalid request body", errors.New("Price must be positive"))
	}

	gallery, err := h.updateGallery(req, galleryId, userId)
	if err != nil {
		return errors.Wrap(err, "failed to update gallery")
	}

	return c.Status(fiber.StatusOK).JSON(dto.HttpResponse[dto.UpdateGalleryResponse]{
		Result: dto.UpdateGalleryResponse{
			ID:          galleryId,
			Name:        gallery.Name,
			Description: gallery.Description,
			Price:       gallery.Price,
		},
	})
}

func (h *Handler) updateGallery(req *dto.UpdateGalleryRequest, galleryId int, userId uint) (*model.Gallery, error) {
	var gallery model.Gallery
	if err := h.store.DB.Transaction(func(tx *gorm.DB) error {
		if err := h.store.DB.First(&gallery, "id = ?", galleryId).Error; err != nil {
			return errors.Wrap(err, "Gallery not found")
		}

		if gallery.PhotographerID != userId {
			return apperror.Forbidden("You are not allowed to update this gallery", nil)
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
