package report

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
)

// @Summary     Get all reports
// @Description Get all reports of a user
// @Tags		customer
// @Router      /api/v1/customers/reports [GET]
// @Security    ApiKeyAuth
// @Param       body  body  dto.CreateReportRequest  true  "Report details"
// @Success     200
// @Failure     400   {object}  dto.HttpError
// @Failure     401   {object}  dto.HttpError
// @Failure     403   {object}  dto.HttpError
// @Failure     404   {object}  dto.HttpError
// @Failure     500   {object}  dto.HttpError
func (h *Handler) HandleGetAllReports(c *fiber.Ctx) error {
	jwtEntity, err := h.authMiddleware.GetJWTEntityFromContext(c.UserContext())
	if err != nil {
		return errors.Wrap(err, "Failed getting jwtEntity from context")
	}

	var req dto.PaginationRequest
	if err := c.QueryParser(&req); err != nil {
		return apperror.BadRequest("Iinvalid query", err)
	}
	if err := h.validate.Struct(req); err != nil {
		return apperror.BadRequest("Invalid parameter", err)
	}

	var reports []dto.ReportResponse

	// photographers cannot get their own reports
	if jwtEntity.Role == model.UserRoleCustomer {
		reports, err = h.getAllReports(jwtEntity.ID)
		if err != nil {
			return errors.Wrap(err, "Failed getting reports")
		}
	} else {
		return apperror.Forbidden("Only customers can get reports", nil)
	}

	return c.Status(fiber.StatusOK).JSON(reports)
}

func (h *Handler) getAllReports(userID uint) ([]dto.ReportResponse, error) {

	var reports []model.Report
	if err := h.store.DB.
		Where("reporter_id = ?", userID).Find(&reports).Error; err != nil {
		return nil, errors.Wrap(err, "Failed getting reports")
	}
	return dto.ToReportResponses(reports), nil
}
