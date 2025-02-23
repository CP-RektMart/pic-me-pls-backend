package dto

import (
	"time"

	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
)

type CreateQuotationRequest struct {
	PackageID   uint    `json:"packageId" validate:"required"`
	CustomerID  uint    `json:"customerId" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
	Description string  `json:"description,omitempty"`

	FromDate time.Time `json:"fromDate" validate:"required"`
	ToDate   time.Time     `json:"toDate" validate:"required"`
}

type CreateQuotationResponse struct {
	QuotationID uint `json:"quotationId"`
}

type UpdateQuotationRequest struct {
	CreateQuotationRequest
	Status string `json:"status" validate:"required"`
}

type QuotationResponse struct {
	ID       uint    `json:"id"`
	Status   string  `json:"status"`
	Customer string  `json:"customer"`
	Price    float64 `json:"price"`
}

type AcceptQuotationRequest struct {
	QuotationID string `params:"id"`
}

func ToQuotationResponses(quotations []model.Quotation) []QuotationResponse {
	var responses []QuotationResponse
	for _, quotation := range quotations {
		responses = append(responses, QuotationResponse{
			ID:       quotation.ID,
			Status:   quotation.Status.String(),
			Price:    quotation.Price,
			Customer: quotation.Customer.Name,
		})
	}
	return responses
}
