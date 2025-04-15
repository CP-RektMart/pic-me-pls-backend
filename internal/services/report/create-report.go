package report

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
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
			QuotationID:  req.QuotationID,
			ReporterID:   userID,
			ReporterRole: req.ReporterRole,
			Status:       "reported",
			Message:      req.Message,
		}

		var targetQuotation model.Quotation
		if err := tx.First(&targetQuotation, req.QuotationID).Error; err != nil {
			return apperror.NotFound("Quotation not found", err)
		}

		//get photographer id from quotation
		if targetQuotation.PhotographerID == 0 {
			return apperror.NotFound("Quotation not found", errors.New("PhotographerID not found"))
		}

		photographerID := targetQuotation.PhotographerID
		customerID := targetQuotation.CustomerID

		// there is no photographer id in quotation
		if photographerID == 0 {
			return apperror.NotFound("Quotation not found", errors.New("PhotographerID not found"))
		}
		if customerID == 0 {
			return apperror.NotFound("Quotation not found", errors.New("CustomerID not found"))
		}

		// user is not related to the quotation
		if newReport.ReporterRole != "ADMIN" && userID != customerID {
			return apperror.Forbidden("You are not allowed to create a report for this quotation", errors.New("User is not customer or photographer"))
		}

		if err := tx.Create(&newReport).Error; err != nil {
			return errors.Wrap(err, "Failed to create report")
		}

		if err := tx.
			Preload("Quotation").
			First(&newReport, newReport.ID).Error; err != nil {
			return errors.Wrap(err, "Failed fetch report")
		}

		return nil
	}); err != nil {
		return errors.Wrap(err, "Failed to create quotation")
	}

	return nil
}
