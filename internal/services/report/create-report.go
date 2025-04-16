package report

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// @Summary     Create a report
// @Description Creates a new report for a user
// @Tags		customer
// @Router      /api/v1/customer/reports [POST]
// @Security    ApiKeyAuth
// @Param       body  body  dto.CreateReportRequest  true  "Report details"
// @Success     200
// @Failure     400   {object}  dto.HttpError
// @Failure     401   {object}  dto.HttpError
// @Failure     403   {object}  dto.HttpError
// @Failure     404   {object}  dto.HttpError
// @Failure     500   {object}  dto.HttpError
func (h *Handler) HandleCreateReport(c *fiber.Ctx) error {
	userID, err := h.authMiddleware.GetUserIDFromContext(c.UserContext())
	if err != nil {
		return errors.Wrap(err, "Failed to get user id from context")
	}

	req := new(dto.CreateReportRequest)
	if err := c.BodyParser(req); err != nil {
		return apperror.BadRequest("Invalid request body", err)
	}

	if err := h.validate.Struct(req); err != nil {
		return apperror.BadRequest("Invalid request body", err)
	}

	if err := h.createReport(req, userID); err != nil {
		return errors.Wrap(err, "Failed to create report")
	}

	return c.SendStatus(fiber.StatusOK)
}

func (h *Handler) createReport(req *dto.CreateReportRequest, userID uint) error {
	if err := h.store.DB.Transaction(func(tx *gorm.DB) error {
		newReport := model.Report{
			QuotationID: req.QuotationID,
			ReporterID:  userID,
			Status:      model.ReportStatusReported,
			Message:     req.Message,
			Title:       req.Title,
		}

		var targetQuotation model.Quotation
		if err := tx.First(&targetQuotation, req.QuotationID).Error; err != nil {
			return apperror.NotFound("Quotation not found", err)
		}

		customerID := targetQuotation.CustomerID

		// quotation status is PENDING or COMFIRMED -> not allowed to report
		if targetQuotation.Status == "PENDING" || targetQuotation.Status == "CONFIRMED" {
			return apperror.Forbidden("You are not allowed to report this quotation", errors.New("Quotation is not allowed to be reported"))
		}

		// user is not related to the quotation
		if userID != customerID {
			return apperror.Forbidden("You are not allowed to create a report for this quotation", errors.New("User is not allowed to report this quotation"))
		}

		if err := tx.Create(&newReport).Error; err != nil {
			return errors.Wrap(err, "Failed to create report")
		}

		return nil
	}); err != nil {
		return errors.Wrap(err, "Failed to create quotation")
	}

	return nil
}
