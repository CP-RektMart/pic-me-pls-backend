package auth

import (
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
)

// handlerLogout godoc
// @summary logout the user
// @description Removing their authentication token form cache
// @tags auth
// @security Bearer
// @id logout
// @accept json
// @produce json
// @response 204 "No Content"
// @response 400 {object} dto.HttpResponse "Bad Request"
// @response 500 {object} dto.HttpResponse "Internal Server Error"
// @Router /api/v1/auth/logout [POST]
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
