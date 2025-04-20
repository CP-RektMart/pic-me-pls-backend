package packages

import (
	"testing"

	"github.com/CP-RektMart/pic-me-pls-backend/internal/dto"
)

func TestValidateCreatePackageRequest(t *testing.T) {
	tests := []struct {
		name     string
		request  dto.CreatePackageRequest
		userID   uint
		expected bool
	}{
		{
			name: "TC4: Valid request",
			request: dto.CreatePackageRequest{
				Name:        "Test Package",
				Description: "Test Description",
				Price:       100.0,
				CategoryID:  func(u uint) *uint { return &u }(1),
				Media: []dto.MediaPackageRequest{
					{PictureURL: "http://example.com/image.jpg",
						Description: "Test Image"},
				},
			},
			userID:   1,
			expected: true,
		},
		{
			name: "TC5: Invalid request - empty name",
			request: dto.CreatePackageRequest{
				Name:        "",
				Description: "Test Description",
				Price:       100.0,
				CategoryID:  func(u uint) *uint { return &u }(1),
				Media: []dto.MediaPackageRequest{
					{PictureURL: "http://example.com/image.jpg",
						Description: "Test Image"},
				},
			},
			userID:   1,
			expected: false,
		},
		{
			name: "TC6: Invalid request - empty description",
			request: dto.CreatePackageRequest{
				Name:        "Test Package",
				Description: "",
				Price:       100.0,
				CategoryID:  func(u uint) *uint { return &u }(1),
				Media: []dto.MediaPackageRequest{
					{PictureURL: "http://example.com/image.jpg",
						Description: "Test Image"},
				},
			},
			userID:   1,
			expected: false,
		},
		{
			name: "TC7: Invalid request - non positive price",
			request: dto.CreatePackageRequest{
				Name:        "Test Package",
				Description: "Test Description",
				Price:       -100.0,
				CategoryID:  func(u uint) *uint { return &u }(1),
				Media: []dto.MediaPackageRequest{
					{PictureURL: "http://example.com/image.jpg",
						Description: "Test Image"},
				},
			},
			userID:   1,
			expected: false,
		},
		{
			name: "TC8: Invalid request - non positive user ID",
			request: dto.CreatePackageRequest{
				Name:        "Test Package",
				Description: "Test Description",
				Price:       100.0,
				CategoryID:  func(u uint) *uint { return &u }(1),
				Media: []dto.MediaPackageRequest{
					{PictureURL: "http://example.com/image.jpg",
						Description: "Test Image"},
				},
			},
			userID:   0,
			expected: false,
		},
		{
			name: "TC9: Invalid request - empty media",
			request: dto.CreatePackageRequest{
				Name:        "Test Package",
				Description: "Test Description",
				Price:       100.0,
				CategoryID:  func(u uint) *uint { return &u }(1),
				Media:       []dto.MediaPackageRequest{},
			},
			userID:   1,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateCreatePackageRequest(&tt.request, tt.userID)
			if (err == nil) != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, err == nil)
			}
		})
	}
}
