package quotation

import (
	"encoding/json"

	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// @Summary     Create a quotation
// @Description Creates a new quotation for a customer and package
// @Tags        quotations
// @Router      /api/v1/photographer/quotations [POST]
// @Security    ApiKeyAuth
// @Param       body  body  dto.CreateQuotationRequest  true  "Quotation details"
// @Success     204
// @Failure     400   {object}  dto.HttpError
// @Failure     500   {object}  dto.HttpError
func (h *Handler) HandleCreateQuotation(c *fiber.Ctx) error {
	userID, err := h.authMiddleware.GetUserIDFromContext(c.UserContext())
	if err != nil {
		return errors.Wrap(err, "failed to get user id from context")
	}

	req := new(dto.CreateQuotationRequest)
	if err := c.BodyParser(req); err != nil {
		return apperror.BadRequest("invalid request body", err)
	}

	if err := h.validate.Struct(req); err != nil {
		return apperror.BadRequest("invalid request body", err)
	}

	if err := h.CreateQuotation(req, userID); err != nil {
		return errors.Wrap(err, "failed to create quotation")
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *Handler) CreateQuotation(req *dto.CreateQuotationRequest, photographerID uint) error {

	// Validate the request
	if err := ValidateCreateQuotationRequest(req, photographerID); err != nil {
		return errors.Wrap(err, "Invalid request")
	}

	if err := h.store.DB.Transaction(func(tx *gorm.DB) error {
		newQuotation := model.Quotation{
			CustomerID:     req.CustomerID,
			PackageID:      req.PackageID,
			Description:    req.Description,
			Price:          req.Price,
			FromDate:       req.FromDate,
			ToDate:         req.ToDate,
			Status:         model.QuotationPending,
			PhotographerID: photographerID,
		}

		var customer model.User
		if err := tx.First(&customer, req.CustomerID).Error; err != nil {
			return apperror.NotFound("customer not found", err)
		}

		var targetPackage model.Package
		if err := tx.First(&targetPackage, req.PackageID).Error; err != nil {
			return apperror.NotFound("package not found", err)
		}

		var photographer model.Photographer
		if err := tx.First(&photographer, photographerID).Error; err != nil {
			return apperror.NotFound("Photographer not found", err)
		}

		if photographer.IsBanned {
			return apperror.Forbidden("You are banned from creating quotations", errors.New("banned photographer"))
		}

		if err := tx.Create(&newQuotation).Error; err != nil {
			return errors.Wrap(err, "failed to create quotation")
		}

		if err := tx.
			Preload("Package.Category").
			Preload("Package.Tags").
			First(&newQuotation, newQuotation.ID).Error; err != nil {
			return errors.Wrap(err, "failed fetch quotation")
		}

		json, err := json.Marshal(dto.ToQuotationMessageResponse(newQuotation))
		if err != nil {
			return errors.Wrap(err, "failed Marshal quotation json")
		}

		if err := h.chatService.SendMessageModel(model.Message{
			Type:       model.MessageTypeQuotation,
			Content:    string(json),
			SenderID:   newQuotation.PhotographerID,
			ReceiverID: newQuotation.CustomerID,
		}); err != nil {
			return errors.Wrap(err, "failed send message")
		}

		return nil
	}); err != nil {
		return errors.Wrap(err, "failed to create quotation")
	}

	return nil
}

func ValidateCreateQuotationRequest(req *dto.CreateQuotationRequest, photographerID uint) error {
	if req.CustomerID <= 0 {
		return apperror.BadRequest("Customer ID must be positive", nil)
	}
	if req.PackageID <= 0 {
		return apperror.BadRequest("Package ID must be postive", nil)
	}
	if photographerID <= 0 {
		return apperror.BadRequest("Photographer ID must be positive", nil)
	}
	if req.FromDate.IsZero() {
		return apperror.BadRequest("From date is required", nil)
	}
	if req.ToDate.IsZero() {
		return apperror.BadRequest("To date is required", nil)
	}
	if req.Price <= 0 {
		return apperror.BadRequest("Price must be positive", nil)
	}
	if req.Description == "" {
		return apperror.BadRequest("Description is required", nil)
	}
	return nil
}
