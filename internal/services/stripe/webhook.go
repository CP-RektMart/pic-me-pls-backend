package stripe

import (
	"fmt"
	"log"
	"os"

	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/gofiber/fiber/v2"
	"github.com/stripe/stripe-go/v78/webhook"
)

// HandleStripeWebhook updates quotation status when payment is successful
func (h *Handler) HandleStripeWebhook(c *fiber.Ctx) error {
	// Read the request body
	payload := c.Body()

	// Get Stripe signature from headers
	signatureHeader := c.Get("Stripe-Signature")
	secret := os.Getenv("STRIPE_WEBHOOK_SECRET")

	// Verify webhook signature
	event, err := webhook.ConstructEventWithOptions(
		payload,
		signatureHeader,
		secret,
		webhook.ConstructEventOptions{
			IgnoreAPIVersionMismatch: true,
		},
	)

	if err != nil {
		return apperror.BadRequest("Invalid webhook signature", err)
	}

	// Handle successful payment
	if event.Type == "checkout.session.completed" {
		clientReferenceID, ok := event.Data.Object["client_reference_id"].(string)
		if !ok {
			return apperror.BadRequest("Missing client_reference_id", nil)
		}

		var quotationID int
		_, err := fmt.Sscanf(clientReferenceID, "%d", &quotationID)
		if err != nil {
			return apperror.BadRequest("Invalid quotation ID", err)
		}

		// Update quotation status in database
		if err := h.store.DB.Model(&model.Quotation{}).
			Where("id = ?", quotationID).
			Update("status", model.QuotationPaid).Error; err != nil {
			return apperror.Internal("Failed to update quotation status", err)
		}

		log.Printf("Quotation %d marked as PAID (Stripe Session: %s)", quotationID, event.Data.Object["id"].(string))
	}

	return c.SendStatus(fiber.StatusOK)
}
