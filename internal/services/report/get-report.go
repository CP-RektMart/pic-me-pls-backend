package report

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
)

// @Summary     Get report by id
// @Description Get a report of a user by id
// @Tags		customer
// @Router      /api/v1/customer/reports/{id} [GET]
// @Security    ApiKeyAuth
// @Param       body  body  dto.CreateReportRequest  true  "Report details"
// @Success     200
// @Failure     400   {object}  dto.HttpError
// @Failure     401   {object}  dto.HttpError
// @Failure     403   {object}  dto.HttpError
// @Failure     404   {object}  dto.HttpError
// @Failure     500   {object}  dto.HttpError
func (h *Handler) HandleGetReportByID(c *fiber.Ctx) error {

	userID, err := h.authMiddleware.GetUserIDFromContext(c.UserContext())
	if err != nil {
		return errors.Wrap(err, "Failed to get user id from context")
	}

	req := new(dto.GetReportByIDRequest)
	if err := c.ParamsParser(req); err != nil {
		return apperror.BadRequest("Invalid params", err)
	}

	report, err := h.getReport(req.ReportID, userID)
	if err != nil {
		return errors.Wrap(err, "Failed to get report")
	}

	return c.Status(fiber.StatusOK).JSON(dto.ReportResponse{
		ReportID:    report.ID,
		QuotationID: report.QuotationID,
		ReporterID:  report.ReporterID,
		Status:      string(report.Status),
		Message:     report.Message,
		Title:       report.Title,
	})
}

func (h *Handler) getReport(id uint, userID uint) (*model.Report, error) {
	var report model.Report
	if err := h.store.DB.Where("id = ? AND reporter_id = ?", id, userID).First(&report).Error; err != nil {
		return nil, apperror.NotFound("Report not found", err)
	}

	return &report, nil
}
