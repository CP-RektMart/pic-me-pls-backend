package dto

import (
	"time"

	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
)

type AcceptQuotationRequest struct {
	QuotationID string `params:"id"`
}

type GetQuotationRequest struct {
	QuotationID string `params:"id"`
}

type ConfirmQuotationRequest struct {
	QuotationID string `params:"id"`
}

type CancelQuotationRequest struct {
	QuotationID string `params:"id"`
}

type UpdateQuotationRequest struct {
	QuotationID string `params:"id"`
}

type CreateQuotationRequest struct {
	PackageID   uint    `json:"packageId" validate:"required"`
	CustomerID  uint    `json:"customerId" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
	Description string  `json:"description,omitempty"`

	FromDate time.Time `json:"fromDate" validate:"required"`
	ToDate   time.Time `json:"toDate" validate:"required"`
}

type QuotationResponse struct {
	ID           uint                  `json:"id"`
	Package      PackageResponse       `json:"package,omitempty"`
	Customer     UserResponse          `json:"customer,omitempty"`
	Photographer PhotographerResponse  `json:"photographer,omitempty"`
	Status       model.QuotationStatus `json:"status"`
	Price        float64               `json:"price"`
	Description  string                `json:"description"`
	FromDate     time.Time             `json:"fromDate"`
	ToDate       time.Time             `json:"toDate"`
}

func ToQuotationResponse(quotation model.Quotation) QuotationResponse {
	return QuotationResponse{
		ID:           quotation.ID,
		Package:      ToPackageResponse(quotation.Package),
		Customer:     ToUserResponse(quotation.Customer),
		Photographer: ToPhotographerResponse(quotation.Photographer),
		Status:       quotation.Status,
		Price:        quotation.Price,
		Description:  quotation.Description,
		FromDate:     quotation.FromDate,
		ToDate:       quotation.ToDate,
	}
}

func ToQuotationResponses(quotations []model.Quotation) []QuotationResponse {
	responses := make([]QuotationResponse, 0, len(quotations))
	for _, quotation := range quotations {
		responses = append(responses, ToQuotationResponse(quotation))
	}
	return responses
}
