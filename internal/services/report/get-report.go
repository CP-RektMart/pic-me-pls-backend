package report

import (
	_ "github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/gofiber/fiber/v2"
)

// @Summary			get report by id
// @Tags			reports
// @Router			/api/v1/reports/{reportID} [GET]
// @Security		ApiKeyAuth
// @Param 			reportID 	path 	uint 	true 	"reportID"
// @Success			200 {object} 	dto.HttpResponse[dto.ReportResponse]
// @Failure			400	{object}	dto.HttpError
// @Failure			401	{object}	dto.HttpError
// @Failure			403	{object}	dto.HttpError
// @Failure			404	{object}	dto.HttpError
// @Failure			500	{object}	dto.HttpError
func (h *Handler) HandleGetReportByID(c *fiber.Ctx) error {
	return nil
}
