package category

import (
	"testing"

	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
)

func TestValidateCreateCategoryRequest(t *testing.T) {
	tests := []struct {
		name     string
		request  dto.CreateCategoryRequest
		expected bool
	}{
		{
			name: "TC1: Valid request",
			request: dto.CreateCategoryRequest{
				Name:        "Test Category",
				Description: "Test Description",
			},
			expected: true,
		},
		{
			name: "TC2: Invalid request - empty name",
			request: dto.CreateCategoryRequest{
				Name:        "",
				Description: "Test Description",
			},
			expected: false,
		},
		{
			name: "TC3: Invalid request - empty description",
			request: dto.CreateCategoryRequest{
				Name:        "Test Category",
				Description: "",
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateCreateCategoryRequest(tt.request)
			if (err == nil) != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, err == nil)
			}
		})
	}
}
