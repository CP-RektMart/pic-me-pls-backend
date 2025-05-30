package dto

import (
	"time"

	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/samber/lo"
)

type CompleteQuotationRequest struct {
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

type CreateCheckoutSessionQuotationRequest struct {
	QuotationID string `params:"id"`
}

type UpdateQuotationRequest struct {
	QuotationID string `params:"id" validate:"required"`

	PackageID   uint    `json:"packageId" validate:"required"`
	CustomerID  uint    `json:"customerId" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
	Description string  `json:"description,omitempty"`

	FromDate time.Time `json:"fromDate" validate:"required"`
	ToDate   time.Time `json:"toDate" validate:"required"`
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

type GetQuotationResponse struct {
	ID           uint                  `json:"id"`
	Package      PackageResponse       `json:"package,omitempty"`
	Customer     UserResponse          `json:"customer,omitempty"`
	Photographer PhotographerResponse  `json:"photographer,omitempty"`
	Status       model.QuotationStatus `json:"status"`
	Price        float64               `json:"price"`
	Description  string                `json:"description"`
	FromDate     time.Time             `json:"fromDate"`
	ToDate       time.Time             `json:"toDate"`
	Previews     []ListPreviewResponse `json:"previews,omitempty"`
	Review       *ReviewResponse       `json:"review,omitempty"`
}
type QuotationMessageResponse struct {
	ID          uint                            `json:"id"`
	Package     QuotationMessagePackageResponse `json:"package"`
	Status      model.QuotationStatus           `json:"status"`
	Price       float64                         `json:"price"`
	Description string                          `json:"description"`
	FromDate    time.Time                       `json:"fromDate"`
	ToDate      time.Time                       `json:"toDate"`
}

func ToGetQuotationResponse(quotation model.Quotation) GetQuotationResponse {
	return GetQuotationResponse{
		ID:           quotation.ID,
		Package:      ToPackageResponse(quotation.Package),
		Customer:     ToUserResponse(quotation.Customer),
		Photographer: ToPhotographerResponse(quotation.Photographer),
		Status:       quotation.Status,
		Price:        quotation.Price,
		Description:  quotation.Description,
		FromDate:     quotation.FromDate,
		ToDate:       quotation.ToDate,
		Previews:     ToListPreviewResponses(quotation.Previews),
		Review:       lo.EmptyableToPtr(ToReviewResponse(quotation.Review)),
	}
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
	return lo.Map(quotations, func(quotation model.Quotation, _ int) QuotationResponse {
		return ToQuotationResponse(quotation)
	})
}

func ToQuotationMessageResponse(quotation model.Quotation) QuotationMessageResponse {
	return QuotationMessageResponse{
		ID:          quotation.ID,
		Package:     ToQuotationMessagePackageResponse(quotation.Package),
		Status:      quotation.Status,
		Price:       quotation.Price,
		Description: quotation.Description,
		FromDate:    quotation.FromDate,
		ToDate:      quotation.ToDate,
	}
}

type CheckoutSessionResponse struct {
	CheckoutURL string `json:"checkout_url"`
}
