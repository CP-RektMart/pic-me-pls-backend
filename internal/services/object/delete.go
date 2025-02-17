package object

import (
	"fmt"
	"strings"

	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) Delete(c *fiber.Ctx) error {
	URL, ok := c.Queries()["URL"]
	if !ok {
		return apperror.BadRequest("No URL specify", nil)
	}

	prefix := fmt.Sprintf("%s/storage/v1/object/public/%s/", h.config.Url, h.config.Bucket)
	path := strings.TrimPrefix(URL, prefix)

	if err := h.validatePath(path); err != nil {
		return err
	}

	if err := h.store.Storage.DeleteFile(c.UserContext(), path); err != nil {
		return errors.Wrap(err, "Failed delete file")
	}

	return c.SendStatus(fiber.StatusNoContent)
}
