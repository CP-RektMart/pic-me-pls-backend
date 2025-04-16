package report

import (
	_ "github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/gofiber/fiber/v2"
)

// @Summary			list reports
// @Tags			reports
// @Router			/api/v1/reports [GET]
// @Security		ApiKeyAuth
// @Param        page      query    int    	    false  "Page number"
// @Param        pageSize  query    int    		false  "Page size"
// @Success			200 {object} 	dto.HttpResponse[dto.PaginationResponse[dto.ReportResponse]]
// @Failure			400	{object}	dto.HttpError
// @Failure			401	{object}	dto.HttpError
// @Failure			403	{object}	dto.HttpError
// @Failure			404	{object}	dto.HttpError
// @Failure			500	{object}	dto.HttpError
func (h *Handler) HandleListReports(c *fiber.Ctx) error {
	return nil
}
