package admin

import (
	_ "github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/gofiber/fiber/v2"
)

// @Summary			verify photographer
// @Tags			admin
// @Router			/api/v1/admin/photographers/{photographerID}/verify [PATCH]
// @Security		ApiKeyAuth
// @Param 			photographerID 	path 	uint 	true 	"photographerID"
// @Success			200
// @Failure			400	{object}	dto.HttpError
// @Failure			401	{object}	dto.HttpError
// @Failure			403	{object}	dto.HttpError
// @Failure			404	{object}	dto.HttpError
// @Failure			500	{object}	dto.HttpError
func (h *Handler) HandleVerifyPhotographer(c *fiber.Ctx) error {
	return nil
}
