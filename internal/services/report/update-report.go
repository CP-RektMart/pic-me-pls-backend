package report

import (
	_ "github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/gofiber/fiber/v2"
)

// @Summary			update report
// @Tags			reports
// @Router			/api/v1/reports/{reportID} [PATCH]
// @Security		ApiKeyAuth
// @Param 			reportID 	path 	uint 	true 	"reportID"
// @Param 			RequestBody 	body 	dto.UpdateReportRequest 	true 	"request body"
// @Success			200
// @Failure			400	{object}	dto.HttpError
// @Failure			401	{object}	dto.HttpError
// @Failure			403	{object}	dto.HttpError
// @Failure			404	{object}	dto.HttpError
// @Failure			500	{object}	dto.HttpError
func (h *Handler) HandleUpdateReportRequest(c *fiber.Ctx) error {
	return nil
}
