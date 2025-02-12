package auth

import (
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
)

// @Summary			Logout
// @Description		Logout
// @Tags			auth
// @Router			/api/v1/auth/logout [POST]
// @Security		ApiKeyAuth
// @Success			204
// @Failure			400	{object}	dto.HttpResponse
// @Failure			500	{object}	dto.HttpResponse
func (h *Handler) HandleLogout(c *fiber.Ctx) error {
	ctx := c.UserContext()

	userID, err := h.authmiddleware.GetUserIDFromContext(ctx)
	if err != nil {
		return errors.Wrap(err, "failed to get user")
	}

	if err := h.jwtService.RemoveToken(ctx, userID); err != nil {
		return errors.Wrap(err, "failed to remove token")
	}

	return c.SendStatus(fiber.StatusNoContent)
}
