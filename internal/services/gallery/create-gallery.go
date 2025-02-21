package gallery

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// to be fixed:
// media model ช่วยเพิ่ม description field เข้ามาด้วย
// request body เพิ่ม field media ด้วย จะเป็น list ของข้อมูลที่ใช้สร้าง media
// ช่วยใช้ transaction ด้วยนะ
// ถ้า success แค่คืน status 204 ก็พอ (fiber.StatusNoContent)
// update swagger too

// @Summary			Create gallery
// @Description		Create gallery by photographer
// @Tags			gallery
// @Router			/api/v1/gallery [POST]
// @Security		ApiKeyAuth
// @Accept			multipart/form-data
// @Param			name 			formData	string	true 	"Gallery name"
// @Param			description 	formData 	string 	true 	"Gallery description"
// @Param			price 			formData 	number 	true 	"Gallery price"
// @Param			galleryPhotos 	formData 	file 	true 	"Gallery photos"
// @Success			200	{object}	dto.HttpResponse[dto.CreateGalleryResponse]
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

	return c.Status(fiber.StatusOK).JSON(dto.HttpResponse[dto.CreateGalleryResponse]{
		Result: dto.CreateGalleryResponse{
			ID:               createdGallery.ID,
			Name:             createdGallery.Name,
			Description:      createdGallery.Description,
			Price:            createdGallery.Price,
			PhotographerID:   createdGallery.PhotographerID,
			PhotographerName: createdGallery.Photographer.User.Name,
		},
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
