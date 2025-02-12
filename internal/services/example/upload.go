package example

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

func (h *Handler) HandlerUploadExample(c *fiber.Ctx) error {
	ctx := c.UserContext()

	file, err := c.FormFile("file")
	if err != nil {
		return apperror.BadRequest("failed to get file", err)
	}

	contentType := file.Header.Get("Content-Type")

	src, err := file.Open()
	if err != nil {
		return errors.Wrap(err, "failed to open file")
	}
	defer src.Close()

	signedURL, err := h.store.Storage.UploadFile(ctx, file.Filename, contentType, src, true)
	if err != nil {
		return errors.Wrap(err, "failed to upload file")
	}

	return c.JSON(dto.HttpResponse{
		Result: signedURL,
	})
}
