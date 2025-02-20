package quotation

import (
	"strconv"

	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) HandleGetQuotationByID(c *fiber.Ctx) error {

	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return apperror.BadRequest("invalid id", err)
	}

	var quotation model.Quotation
	if err := h.store.DB.
		Preload("Customer").
		Where("id = ?", id).
		First(&quotation).
		Error; err != nil {
		return apperror.NotFound("quotation not found", err)
	}

	result := dto.ToQuotationResponses([]model.Quotation{quotation})[0]

	return c.JSON(dto.HttpResponse[dto.QuotationResponse]{
		Result: result,
	})

}
