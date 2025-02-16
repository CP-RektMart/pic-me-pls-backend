package gallery

import (
	"context"
	"fmt"
	"mime/multipart"
	"strconv"

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
// @Param 			name 			formData 	string		true	"Gallery name"
// @Param 			description		formData 	string		true	"Description"
// @Param 			price			formData 	int		true	"Price"
// @Param 			galleryPhotos 	formData 	file		true	"Gallery photos"
// @Success			200	{object}	dto.HttpResponse{result=dto.GalleryResponse}
// @Failure			400	{object}	dto.HttpResponse
// @Failure			500	{object}	dto.HttpResponse
func (h *Handler) HandleCreateGallery(c *fiber.Ctx) error {
	userId, err := h.authMiddleware.GetUserIDFromContext(c.UserContext())
	if err != nil {
		return errors.Wrap(err, "failed to get user id from context")
	}

	req := new(dto.GalleryRequest)
	req.Name = c.FormValue("name")
	req.Description = c.FormValue("description")
	price, err := strconv.ParseFloat(c.FormValue("price"), 64)
	if err != nil {
		return apperror.BadRequest("invalid price value", err)
	}
	req.Price = price

	if err := h.validate.Struct(req); err != nil {
		return apperror.BadRequest("invalid request body", err)
	}

	form, err := c.MultipartForm()
	if err != nil {
		return apperror.BadRequest("failed to parse multipart form", err)
	}
	files := form.File["galleryPhotos"]
	if len(files) == 0 {
		return apperror.BadRequest("Gallery picture is required", errors.Errorf("field missing"))
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

	var uploadedPhotoURLs []string
	for _, file := range files {
		fmt.Println("Processing file:", file.Filename)

		signedURL, err := h.uploadGalleryPhoto(c.UserContext(), file, galleryFolder(createdGallery.ID))
		if err != nil {
			return errors.Wrap(err, "failed to upload photos")
		}

		uploadedPhotoURLs = append(uploadedPhotoURLs, signedURL)
	}

	response := dto.GalleryResponse{
		ID:               createdGallery.ID,
		Name:             createdGallery.Name,
		Description:      createdGallery.Description,
		Price:            createdGallery.Price,
		PhotographerID:   createdGallery.PhotographerID,
		PhotographerName: createdGallery.Photographer.User.Name,
		GalleryPhotos:    uploadedPhotoURLs,
	}

	return c.Status(fiber.StatusOK).JSON(dto.HttpResponse{
		Result: response,
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

func (h *Handler) uploadGalleryPhoto(c context.Context, file *multipart.FileHeader, folder string) (string, error) {
	contentType := file.Header.Get("Content-Type")

	src, err := file.Open()
	if err != nil {
		return "", errors.Wrap(err, "failed to open file")
	}
	defer src.Close()

	signedURL, err := h.store.Storage.UploadFile(c, folder+file.Filename, contentType, src, true)
	if err != nil {
		return "", errors.Wrap(err, "failed to upload file")
	}

	return signedURL, nil
}

func galleryFolder(galleryId uint) string {
	return "gallery_photos/" + strconv.FormatUint(uint64(galleryId), 10) + "/"
}
