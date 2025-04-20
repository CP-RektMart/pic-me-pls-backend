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
			name: "TC10: Valid request",
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
			name: "TC11: Invalid request - non positive customer ID",
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
			name: "TC12: Invalid request - non positive package ID",
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
			name: "TC13: Invalid request - non positive price",
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
			name: "TC14: Invalid request - non positive photographer ID",
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
			name: "TC15: Invalid request - from date not provided",
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
			name: "TC16: Invalid request - to date not provided",
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
			name: "TC17: Invalid request - empty description",
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
