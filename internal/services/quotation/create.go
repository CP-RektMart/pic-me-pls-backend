package quotation

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// @Summary     Create a quotation
// @Description Creates a new quotation for a customer and package
// @Tags        quotation
// @Router      /api/v1/quotations [POST]
// @Security    ApiKeyAuth
// @Param       body  body  dto.QuotationRequest  true  "Quotation details"
// @Success     201   {object}  dto.HttpResponse[dto.QuotationResponse]
// @Failure     400   {object}  dto.HttpError
// @Failure     500   {object}  dto.HttpError
func (h *Handler) HandleCreate(c *fiber.Ctx) error {
	userID, err := h.authMiddleware.GetUserIDFromContext(c.UserContext())
	if err != nil {
		return errors.Wrap(err, "failed to get user id from context")
	}

	req := new(dto.QuotationRequest)
	if err := c.BodyParser(req); err != nil {
		return apperror.BadRequest("invalid request body", err)
	}

	if err := h.validate.Struct(req); err != nil {
		return apperror.BadRequest("invalid request body", err)
	}

	newQuotation, err := h.CreateQuotation(req,userID)
	if (err != nil) {
		return errors.Wrap(err, "failed to create quotation")
	}

	return c.Status(fiber.StatusCreated).JSON(dto.HttpResponse[dto.QuotationResponse]{
		Result: dto.QuotationResponse{
			ID: newQuotation.ID,
			Customer: newQuotation.Customer.Name,
			Status: newQuotation.Status.String(),
			Price: newQuotation.Price,
		},
	})
}

func (h *Handler) CreateQuotation(req *dto.QuotationRequest, userID uint) (*model.Quotation, error) {
	var newQuotation model.Quotation

	if err := h.store.DB.Transaction(func(tx *gorm.DB) error {
		newQuotation = model.Quotation{
			CustomerID: req.CustomerID,
			PackageID: req.PackageID,
			Description:    req.Description,
			Price:          req.Price,
			FromDate: req.FromDate,
			ToDate:  req.ToDate,
			Status: model.QuotationPending,
		}

		// Check CustomerID and PackageID existed in database
		var customer model.User
		if err := tx.First(&customer, req.CustomerID).Error; err != nil {
			return apperror.NotFound("customer not found", err)
		}
		newQuotation.Customer = customer

		var targetPackage model.Package
		if err := tx.First(&targetPackage, req.PackageID).Error; err != nil {
			return apperror.NotFound("package not found", err)
		}
		newQuotation.Package = targetPackage

		// Find the photographer associated with the user
		var photographer model.Photographer
		if err := tx.First(&photographer, "user_id = ?", userID).Error; err != nil {
			return apperror.NotFound("Photographer not found for user", err)
		}
		newQuotation.PhotographerID = photographer.ID
		newQuotation.Photographer = photographer

		// Create Quotation  
		if err := tx.Create(&newQuotation).Error; err != nil {
			return errors.Wrap(err, "failed to create quotation")
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return &newQuotation, nil
} 