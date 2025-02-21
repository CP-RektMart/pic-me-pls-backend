package object

import (
	"fmt"
	"strings"

	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// @Summary			Upload image
// @Description			receive formData body, path (string, folder path, don't include ".." or prefix with "/") and file
// @Tags			objects
// @Router			/api/v1/objects [POST]
// @Param 			file 	formData 	file		true	"picture (optional)"
// @Param 			folder 	formData 	string		false	"folder enum (PACKAGE, VERIFY_CITIZENCARD, PROFILE_IMAGE)"
// @Success			200	{object}	dto.HttpResponse[dto.ObjectUploadResponse]
// @Failure			400	{object}	dto.HttpError
// @Failure			500	{object}	dto.HttpError
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

	folder := Folder(c.FormValue("folder"))
	fileName := fmt.Sprintf("%s-%s", uuid.NewString(), file.Filename)
	path := folder.GetFullPath(fileName)
	contentType := file.Header.Get("Content-Type")

	if err := h.validatePath(path); err != nil {
		return err
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

func (h *Handler) validatePath(path string) error {
	if strings.Contains(path, "..") || strings.HasPrefix(path, "/") {
		return apperror.BadRequest("invalid path", nil)
	}
	return nil
}
