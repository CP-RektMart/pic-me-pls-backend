package customer

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// @Summary			Get customer public profile
// @Description			Get customer public profile
// @Tags			customer
// @Router			/api/v1/customers/{id} [GET]
// @Security			ApiKeyAuth
// @Param 			id 	path 	uint 	true 	"customer's userId"
// @Success			200	{object}	dto.HttpResponse[dto.CustomerPublicResponse]
// @Failure			404	{object}	dto.HttpError
// @Failure			500	{object}	dto.HttpError
func (h *Handler) HandlerCustomerPublicProfile(c *fiber.Ctx) error {
	var req dto.GetCustomerPublicRequest
	if err := c.ParamsParser(&req); err != nil {
		return apperror.BadRequest("invalid path parameters", err)
	}
	if err := h.validate.Struct(req); err != nil {
		return apperror.BadRequest("invalid path parameters", err)
	}

	var customer model.User
	if err := h.store.DB.First(&customer, req.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return apperror.NotFound("customer not found", err)
		}
		return errors.Wrap(err, "failed fetch customer")
	}

	return c.Status(fiber.StatusOK).JSON(dto.HttpResponse[dto.CustomerPublicResponse]{
		Result: dto.ToCustomerPublicResponse(customer),
	})
}
