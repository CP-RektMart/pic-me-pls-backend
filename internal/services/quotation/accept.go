package quotation

import (
	"context"

	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/cockroachdb/errors"
	"github.com/danielgtaylor/huma/v2"
	"gorm.io/gorm"
)

func (h *Handler) Accept(ctx context.Context, req *dto.AcceptQuotationRequest) (*struct{}, error) {
	userID, err := h.authMiddleware.GetUserIDFromContext(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed get user id from context")
	}

	var quotation *model.Quotation
	if err := h.store.DB.First(&quotation, req.QuotationID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, huma.Error404NotFound("quotation not found", err)
		}
		return nil, errors.Wrap(err, "failed find quotation")
	}

	if quotation.CustomerID != userID {
		return nil, huma.Error403Forbidden("user not have permission", nil)
	}

	quotation.Status = model.QuotationConfirm
	if err := h.store.DB.Save(&quotation).Error; err != nil {
		return nil, errors.Wrap(err, "failed confirm quotation")
	}

	return nil, nil
}
