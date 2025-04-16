package admin

import (
	_ "github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/gofiber/fiber/v2"
)

// @Summary			get package by id
// @Tags			admin
// @Router			/api/v1/admin/packages/{packageID} [GET]
// @Security		ApiKeyAuth
// @Param 			packageID 	path 	uint 	true 	"packageID"
// @Success			200 {object}	dto.HttpResponse[dto.PackageResponse]
// @Failure			400	{object}	dto.HttpError
// @Failure			401	{object}	dto.HttpError
// @Failure			403	{object}	dto.HttpError
// @Failure			404	{object}	dto.HttpError
// @Failure			500	{object}	dto.HttpError
func (h *Handler) HandleGetPackages(c *fiber.Ctx) error {
	return nil
}
