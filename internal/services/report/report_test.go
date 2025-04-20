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
			name: "TC:18 Valid request",
			request: dto.CreateReportRequest{
				QuotationID: 1,
				Message:     "Test message",
				Title:       "Test title",
			},
			userID:   1,
			expected: true,
		},
		{
			name: "TC19: Invalid request - non positive quotation ID",
			request: dto.CreateReportRequest{
				QuotationID: 0,
				Message:     "Test message",
				Title:       "Test title",
			},
			userID:   1,
			expected: false,
		},
		{
			name: "TC20: Invalid request - empty message",
			request: dto.CreateReportRequest{
				QuotationID: 1,
				Message:     "",
				Title:       "Test title",
			},
			userID:   1,
			expected: false,
		},
		{
			name: "TC21: Invalid request - empty title",
			request: dto.CreateReportRequest{
				QuotationID: 1,
				Message:     "Test message",
				Title:       "",
			},
			userID:   1,
			expected: false,
		},
		{
			name: "TC22: Invalid request - non positive user ID",
			request: dto.CreateReportRequest{
				QuotationID: 1,
				Message:     "Test message",
				Title:       "Test title",
			},
			userID:   0,
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
