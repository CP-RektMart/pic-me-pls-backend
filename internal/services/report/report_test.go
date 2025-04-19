package report

import (
	"testing"

	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
)

func TestValidateReportRequest(t *testing.T) {
	tests := []struct {
		name     string
		request  dto.CreateReportRequest
		userID   uint
		expected bool
	}{
		{
			name: "valid request",
			request: dto.CreateReportRequest{
				QuotationID: 1,
				Message:     "Test message",
				Title:       "Test title",
			},
			userID:   1,
			expected: true,
		},
		{
			name: "invalid request - empty message",
			request: dto.CreateReportRequest{
				QuotationID: 1,
				Message:     "",
				Title:       "Test title",
			},
			userID:   1,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ValidateReportRequest(&tt.request, tt.userID)
			if (result == nil) != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result == nil)
			}
		})
	}
}
