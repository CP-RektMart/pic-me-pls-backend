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
// @Description Creates a new quotation for a customer and gallery
// @Tags        quotation
// @Router      /api/v1/quotations [POST]
// @Security    ApiKeyAuth
// @Param       body  body  dto.CreateQuotationRequest  true  "Quotation details"
// @Success     201   {object}  dto.HttpResponse[dto.CreateQuotationResponse]
// @Failure     400   {object}  dto.HttpError
// @Failure     500   {object}  dto.HttpError
func (h *Handler) HandleCreate(c *fiber.Ctx) error {
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

	QuotationID, err := h.CreateQuotation(req,userID)
	if (err != nil) {
		return errors.Wrap(err, "failed to create quotation")
	}

	return c.Status(fiber.StatusCreated).JSON(dto.HttpResponse[dto.QuotationResponse]{
		Result: dto.QuotationResponse{
			QuotationID: QuotationID,
		},
	})
}

func (h *Handler) CreateQuotation(req *dto.CreateQuotationRequest, userID uint) (uint, error) {
	var quotationID uint
	
	if err := h.store.DB.Transaction(func(tx *gorm.DB) error {
		quotation := &model.Quotation{
			CustomerID: req.CustomerID,
			GalleryID: req.GalleryID,
			Description:    req.Description,
			Price:          req.Price,
			FromDate: req.FromDate,
			ToDate:  req.ToDate,
			Status: model.QuotationPending,
		}

		// Check CustomerID and GalleryID existed in database
		var customer model.User
		if err := tx.First(&customer, req.CustomerID).Error; err != nil {
			return errors.Wrap(err, "customer not found")
		}
		var gallery model.Gallery
		if err := tx.First(&gallery, req.GalleryID).Error; err != nil {
			return errors.Wrap(err, "gallery not found")
		}

		// Find the photographer associated with the user
		var photographer model.Photographer
		if err := tx.First(&photographer, "user_id = ?", userID).Error; err != nil {
			return errors.Wrap(err, "Photographer not found for user")
		}

		quotation.PhotographerID = photographer.ID;

		// Create Quotation  
		if err := tx.Create(&quotation).Error; err != nil {
			return errors.Wrap(err, "failed to create quotation")
		}
		quotationID = quotation.ID
		return nil
	}); err != nil {
		return 0, errors.Wrap(err, "failed to create quotation")
	}
	return quotationID, nil
} 