package auth

import (
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) HandleLogout(c *fiber.Ctx) error {

	ctx := c.UserContext()

	userID, err := h.authmiddleware.GetUserIDFromContext(ctx)
	if err != nil {
		return apperror.UnAuthorized("falied to get user", err)
	}

	if err := h.jwtService.RemoveToken(c.Context(), userID); err != nil {
		return errors.Wrap(err, "failed to remove token")
	}

	return c.SendStatus(fiber.StatusNoContent)

}
