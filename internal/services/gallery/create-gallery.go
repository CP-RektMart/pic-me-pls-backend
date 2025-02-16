package gallery

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func (h *Handler) HandleCreateGallery(c *fiber.Ctx) error {
	userId, err := h.authMiddleware.GetUserIDFromContext(c.UserContext())
	if err != nil {
		return errors.Wrap(err, "failed to get user id from context")
	}

	req := new(dto.GalleryRequest)
	if err := c.BodyParser(req); err != nil {
		return apperror.BadRequest("invalid request body", err)
	}

	if err := h.validate.Struct(req); err != nil {
		return apperror.BadRequest("invalid request body", err)
	}

	gallery := &model.Gallery{
		PhotographerID: userId,
		Name:           req.Name,
		Description:    req.Description,
		Price:          req.Price,
	}

	createdGallery, err := h.CreateGallery(gallery, userId)
	if err != nil {
		return errors.Wrap(err, "failed to create gallery")
	}

	return c.Status(fiber.StatusOK).JSON(dto.HttpResponse{
		Result: createdGallery,
	})
}

func (h *Handler) CreateGallery(gallery *model.Gallery, userId uint) (*model.Gallery, error) {
	if err := h.store.DB.Transaction(func(tx *gorm.DB) error {
		var photographer model.Photographer
		if err := h.store.DB.Preload("User").First(&photographer, "user_id = ?", userId).Error; err != nil {
			return errors.Wrap(err, "Photographer not found")
		}

		gallery.Photographer = photographer

		if err := tx.Create(&gallery).Error; err != nil {
			return errors.Wrap(err, "failed to create gallery")
		}

		return nil
	}); err != nil {
		return nil, errors.Wrap(err, "failed to create gallery")
	}

	return gallery, nil
}
