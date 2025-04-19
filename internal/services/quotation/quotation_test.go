package quotation

import (
	"testing"
	"time"

	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
)

func TestValidateCreateQuotationRequest(t *testing.T) {
	tests := []struct {
		name           string
		request        dto.CreateQuotationRequest
		photographerID uint
		expected       bool
	}{
		{
			name: "valid request",
			request: dto.CreateQuotationRequest{
				CustomerID:  1,
				PackageID:   1,
				Price:       100.0,
				Description: "Test Description",
				FromDate: func() time.Time {
					t, _ := time.Parse("2006-01-02", "2023-10-01")
					return t
				}(),
				ToDate: func() time.Time {
					t, _ := time.Parse("2006-01-02", "2023-10-02")
					return t
				}(),
			},
			photographerID: 1,
			expected:       true,
		},
		{
			name: "invalid request - non positive customer ID",
			request: dto.CreateQuotationRequest{
				CustomerID:  0,
				PackageID:   1,
				Price:       100.0,
				Description: "Test Description",
				FromDate: func() time.Time {
					t, _ := time.Parse("2006-01-02", "2023-10-01")
					return t
				}(),
				ToDate: func() time.Time {
					t, _ := time.Parse("2006-01-02", "2023-10-02")
					return t
				}(),
			},
			photographerID: 1,
			expected:       false,
		},
		{
			name: "invalid request - non positive package ID",
			request: dto.CreateQuotationRequest{
				CustomerID:  1,
				PackageID:   0,
				Price:       100.0,
				Description: "Test Description",
				FromDate: func() time.Time {
					t, _ := time.Parse("2006-01-02", "2023-10-01")
					return t
				}(),
				ToDate: func() time.Time {
					t, _ := time.Parse("2006-01-02", "2023-10-02")
					return t
				}(),
			},
			photographerID: 1,
			expected:       false,
		},
		{
			name: "invalid request - non positive price",
			request: dto.CreateQuotationRequest{
				CustomerID:  1,
				PackageID:   1,
				Price:       -100.0,
				Description: "Test Description",
				FromDate: func() time.Time {
					t, _ := time.Parse("2006-01-02", "2023-10-01")
					return t
				}(),
				ToDate: func() time.Time {
					t, _ := time.Parse("2006-01-02", "2023-10-02")
					return t
				}(),
			},
			photographerID: 1,
			expected:       false,
		},
		{
			name: "invalid request - non positive photographer ID",
			request: dto.CreateQuotationRequest{
				CustomerID:  1,
				PackageID:   1,
				Price:       100.0,
				Description: "Test Description",
				FromDate: func() time.Time {
					t, _ := time.Parse("2006-01-02", "2023-10-01")
					return t
				}(),
				ToDate: func() time.Time {
					t, _ := time.Parse("2006-01-02", "2023-10-02")
					return t
				}(),
			},
			photographerID: 0,
			expected:       false,
		},
		{
			name: "invalid request - from date not provided",
			request: dto.CreateQuotationRequest{
				CustomerID:  1,
				PackageID:   1,
				Price:       100.0,
				Description: "Test Description",
				FromDate:    time.Time{},
				ToDate: func() time.Time {
					t, _ := time.Parse("2006-01-02", "2023-10-02")
					return t
				}(),
			},
			photographerID: 1,
			expected:       false,
		},
		{
			name: "invalid request - to date not provided",
			request: dto.CreateQuotationRequest{
				CustomerID:  1,
				PackageID:   1,
				Price:       100.0,
				Description: "Test Description",
				FromDate: func() time.Time {
					t, _ := time.Parse("2006-01-02", "2023-10-01")
					return t
				}(),
				ToDate: time.Time{},
			},
			photographerID: 1,
			expected:       false,
		},
		{
			name: "invalid request - empty description",
			request: dto.CreateQuotationRequest{
				CustomerID:  1,
				PackageID:   1,
				Price:       100.0,
				Description: "",
				FromDate: func() time.Time {
					t, _ := time.Parse("2006-01-02", "2023-10-01")
					return t
				}(),
				ToDate: func() time.Time {
					t, _ := time.Parse("2006-01-02", "2023-10-02")
					return t
				}(),
			},
			photographerID: 1,
			expected:       false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateCreateQuotationRequest(&tt.request, tt.photographerID)
			if (err == nil) != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, err == nil)
			}
		})
	}
}
