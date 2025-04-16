package report

import (
	_ "github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/gofiber/fiber/v2"
)

// @Summary			create report
// @Tags			reports
// @Router			/api/v1/reports [POST]
// @Security		ApiKeyAuth
// @Param 			RequestBody 	body 	dto.CreateReportRequest 	true 	"request body"
// @Success			200
// @Failure			400	{object}	dto.HttpError
// @Failure			401	{object}	dto.HttpError
// @Failure			403	{object}	dto.HttpError
// @Failure			404	{object}	dto.HttpError
// @Failure			500	{object}	dto.HttpError
func (h *Handler) HandleGetAdmin(c *fiber.Ctx) error {
	return nil
}
