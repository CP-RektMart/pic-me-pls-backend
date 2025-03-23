package stripe

import (
	"fmt"

	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/cockroachdb/errors"
	"github.com/gofiber/fiber/v2"
	"github.com/stripe/stripe-go/v78"
	"github.com/stripe/stripe-go/v78/checkout/session"
	"gorm.io/gorm"
)

// @Summary      Create Stripe Checkout Session
// @Description  Generates a Stripe checkout session for a quotation
// @Tags         stripe
// @Router       /api/v1/stripe/checkout/{id} [post]
// @Param        id   path  int  true  "Quotation ID"
// @Success      200  {object}  dto.CheckoutSessionResponse
// @Failure      400  {object}  dto.HttpError
// @Failure      500  {object}  dto.HttpError
func (h *Handler) HandleCreateCheckoutSession(c *fiber.Ctx) error {
	var req dto.CreateCheckoutSessionQuotationRequest
	if err := c.ParamsParser(&req); err != nil {
		return apperror.BadRequest("Invalid request parameters", err)
	}

	userID, err := h.authMiddleware.GetUserIDFromContext(c.UserContext())
	if err != nil {
		return errors.Wrap(err, "Failed to get user ID from context")
	}

	var quotation model.Quotation
	if err := h.store.DB.Preload("Customer").Preload("Package").
		First(&quotation, req.QuotationID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return apperror.NotFound("Quotation not found", err)
		}
		return errors.Wrap(err, "Error retrieving quotation")
	}

	if quotation.CustomerID != userID {
		return apperror.Forbidden("You can only create a checkout session for your own quotations", nil)
	}

	if quotation.Status != model.QuotationConfirm {
		return apperror.BadRequest("Quotation is not in confirm status", nil)
	}

	stripe.Key = h.stripeConfig.SecretKey
	frontEndUrl := h.frontendUrl

	params := &stripe.CheckoutSessionParams{
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
					Currency: stripe.String("thb"),
					ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
						Name: stripe.String(quotation.Package.Name),
					},
					UnitAmount: stripe.Int64(int64(quotation.Price * 100)),
				},
				Quantity: stripe.Int64(1),
			},
		},
		Mode:              stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL:        stripe.String(fmt.Sprintf("%s/quotation/%d?payment=success", frontEndUrl, quotation.ID)),
		CancelURL:         stripe.String(fmt.Sprintf("%s/quotation/%d?payment=cancel", frontEndUrl, quotation.ID)),
		ClientReferenceID: stripe.String(fmt.Sprintf("%d", quotation.ID)),
	}

	session, err := session.New(params)
	if err != nil {
		return apperror.Internal("Failed to create Stripe session", err)
	}

	return c.Status(fiber.StatusOK).JSON(dto.CheckoutSessionResponse{
		CheckoutURL: session.URL,
	})
}
