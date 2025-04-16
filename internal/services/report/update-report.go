package report

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// @Summary     Update a report
// @Description Updates an existing report
// @Tags		customer
// @Router      /api/v1/customers/reports/{id} [PATCH]
// @Security    ApiKeyAuth
// @Param       body  body  dto.CreateReportRequest  true  "Report details"
// @Success     200
// @Failure     400   {object}  dto.HttpError
// @Failure     401   {object}  dto.HttpError
// @Failure     403   {object}  dto.HttpError
// @Failure     404   {object}  dto.HttpError
// @Failure     500   {object}  dto.HttpError
func (h *Handler) HandleUpdateReport(c *fiber.Ctx) error {
	userID, err := h.authMiddleware.GetUserIDFromContext(c.UserContext())
	if err != nil {
		return errors.Wrap(err, "Failed to get user id from context")
	}

	req := new(dto.UpdateReportRequest)

	if err := c.ParamsParser(req); err != nil {
		return apperror.BadRequest("Invalid params", err)
	}

	if err := c.BodyParser(req); err != nil {
		return apperror.BadRequest("Invalid request body", err)
	}

	if err := h.validate.Struct(req); err != nil {
		return apperror.BadRequest("Invalid request body", err)
	}

	if err := h.updateReport(req, userID); err != nil {
		return errors.Wrap(err, "Failed to update report")
	}

	return c.SendStatus(fiber.StatusOK)
}

func (h *Handler) updateReport(req *dto.UpdateReportRequest, userID uint) error {
	if err := h.store.DB.Transaction(func(tx *gorm.DB) error {
		var report model.Report
		if err := tx.First(&report, req.ReportID).Error; err != nil {
			return apperror.NotFound("Report not found", err)
		}

		if report.ReporterID != userID {
			return apperror.Forbidden("You are not allowed to update this report", nil)
		}

		if err := tx.Model(&report).Where("id = ?", req.ReportID).Updates(model.Report{
			Status:  req.Status,
			Message: req.Message,
			Title:   req.Title,
		}).Error; err != nil {
			return errors.Wrap(err, "failed to update quotation")
		}

		return nil
	}); err != nil {
		return errors.Wrap(err, "Failed to update report")
	}

	return nil
}
