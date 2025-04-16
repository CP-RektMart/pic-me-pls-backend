package admin

import (
	_ "github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/gofiber/fiber/v2"
)

// @Summary			update user role
// @Tags			admin
// @Router			/api/v1/admin/role/users/{userID} [PATCH]
// @Security		ApiKeyAuth
// @Param 			userID 	path 	uint 	true 	"userID"
// @Param 			admin 	query 	bool 	true 	"isAdmin"
// @Success			200
// @Failure			400	{object}	dto.HttpError
// @Failure			401	{object}	dto.HttpError
// @Failure			403	{object}	dto.HttpError
// @Failure			404	{object}	dto.HttpError
// @Failure			500	{object}	dto.HttpError
func (h *Handler) HandleUpdateUserRole(c *fiber.Ctx) error {
	return nil
}
