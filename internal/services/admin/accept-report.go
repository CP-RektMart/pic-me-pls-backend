package admin

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// @Summary			admin accept report
// @Tags			admin
// @Router			/api/v1/admin/reports/{reportID}/accept [PATCH]
// @Security		ApiKeyAuth
// @Param			reportID		path		int	true	"report id"
// @Success			204
// @Failure			400	{object}	dto.HttpError
// @Failure			401	{object}	dto.HttpError
// @Failure			403	{object}	dto.HttpError
// @Failure			404	{object}	dto.HttpError
// @Failure			500	{object}	dto.HttpError
func (h *Handler) HandleAcceptReport(c *fiber.Ctx) error {
	var req dto.AcceptReportRequest
	if err := c.ParamsParser(&req); err != nil {
		return apperror.BadRequest("invalid request", err)
	}
	if err := h.validate.Struct(req); err != nil {
		return apperror.BadRequest("invalid request", err)
	}

	if err := h.acceptReport(req.ReportID); err != nil {
		return errors.Wrap(err, "failed update report status to ACCEPTED")
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *Handler) acceptReport(reportID uint) error {
	var report model.Report
	if err := h.store.DB.First(&report, reportID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return apperror.NotFound("report not found", err)
		}
		return err
	}

	if report.Status == model.ReportStatusAccepted {
		return apperror.BadRequest("report is already accepted", nil)
	}

	report.Status = model.ReportStatusAccepted
	return h.store.DB.Save(&report).Error
}
