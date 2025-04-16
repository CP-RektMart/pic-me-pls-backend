package report

import (
	_ "github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/gofiber/fiber/v2"
)

// @Summary			delete report
// @Tags			reports
// @Router			/api/v1/reports/{reportID} [DELETE]
// @Security		ApiKeyAuth
// @Param 			reportID 	path 	uint 	true 	"reportID"
// @Success			200
// @Failure			400	{object}	dto.HttpError
// @Failure			401	{object}	dto.HttpError
// @Failure			403	{object}	dto.HttpError
// @Failure			404	{object}	dto.HttpError
// @Failure			500	{object}	dto.HttpError
func (h *Handler) HandleDeleteReport(c *fiber.Ctx) error {
	return nil
}
