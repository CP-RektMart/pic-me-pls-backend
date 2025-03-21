package stripe

import (
	"fmt"
	"log"
	"os"

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
	// Parse quotation ID from URL params
	var req dto.GetQuotationRequest
	if err := c.ParamsParser(&req); err != nil {
		return apperror.BadRequest("Invalid request parameters", err)
	}

	// Retrieve quotation from DB
	var quotation model.Quotation
	if err := h.store.DB.Preload("Customer").Preload("Package").
		Where("id = ?", req.QuotationID).
		First(&quotation).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return apperror.NotFound("Quotation not found", err)
		}
		return errors.Wrap(err, "Error retrieving quotation")
	}

	// Ensure quotation is in "CONFIRM" status
	if quotation.Status != model.QuotationConfirm {
		return apperror.BadRequest("Quotation is not in confirm status", nil)
	}

	// Stripe configuration
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")

	// Create Stripe checkout session
	params := &stripe.CheckoutSessionParams{
		PaymentMethodTypes: stripe.StringSlice([]string{"card"}),
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
		SuccessURL:        stripe.String(fmt.Sprintf("http://localhost:5000/id=%d", quotation.ID)),
		CancelURL:         stripe.String("http://localhost:5000/cancel"),
		ClientReferenceID: stripe.String(fmt.Sprintf("%d", quotation.ID)),
	}

	session, err := session.New(params)
	if err != nil {
		return apperror.Internal("Failed to create Stripe session", err)
	}

	log.Printf("Stripe session created: %s\n", session.ID)

	// Return checkout session response
	return c.Status(fiber.StatusOK).JSON(dto.CheckoutSessionResponse{
		CheckoutURL: session.URL,
	})
}
