package object

import (
	"fmt"
	"strings"

	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) Upload(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return apperror.BadRequest("invalid body", err)
	}

	reader, err := file.Open()
	if err != nil {
		return errors.Wrap(err, "failed open file")
	}
	defer reader.Close()

	folder := c.FormValue("path", "others")
	path := fmt.Sprintf("%s/%s", folder, file.Filename)
	contentType := file.Header.Get("Content-Type")

	if strings.Contains(path, "..") || strings.HasPrefix(path, "/") {
		return apperror.BadRequest("invalid path", nil)
	}

	URL, err := h.store.Storage.UploadFile(c.UserContext(), path, contentType, reader, true)
	if err != nil {
		return errors.Wrap(err, "Failed upload file")
	}

	return c.Status(fiber.StatusOK).JSON(dto.HttpResponse[dto.ObjectUploadResponse]{
		Result: dto.ObjectUploadResponse{
			URL: URL,
		},
	})
}
