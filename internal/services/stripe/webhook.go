package stripe

import (
	"fmt"
	"os"

	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/apperror"
	"github.com/CP-RektMart/pic-me-pls-backend/pkg/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/stripe/stripe-go/v78/webhook"
)

func (h *Handler) HandleStripeWebhook(c *fiber.Ctx) error {
	payload := c.Body()
	signatureHeader := c.Get("Stripe-Signature")
	secret := os.Getenv("STRIPE_WEBHOOK_SECRET")

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

	if event.Type == "checkout.session.completed" {
		clientReferenceID, ok := event.Data.Object["client_reference_id"].(string)
		if !ok {
			return apperror.BadRequest("Missing client_reference_id", nil)
		}

		var quotationID int
		_, err := fmt.Sscanf(clientReferenceID, "%d", &quotationID)
		if err != nil {
			return apperror.BadRequest("Invalid clientReference ID", err)
		}

		if err := h.store.DB.Model(&model.Quotation{}).
			Where("id = ?", quotationID).
			Update("status", model.QuotationPaid).Error; err != nil {
			return apperror.Internal("Failed to update quotation status", err)
		}

		logger.Info("Quotation marked as PAID", "quotationID", quotationID, "stripeSessionID", event.Data.Object["id"].(string))
	}

	return c.SendStatus(fiber.StatusOK)
}
