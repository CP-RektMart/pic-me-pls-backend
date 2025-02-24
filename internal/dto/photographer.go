package dto

import (
	"time"

	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
)

// Form data
type CitizenCardRequest struct {
	CitizenID  string    `validate:"required"`
	LaserID    string    `validate:"required"`
	ExpireDate time.Time `validate:"required"`
}

type CitizenCardResponse struct {
	CitizenID  string    `json:"citizenId"`
	LaserID    string    `json:"laserId"`
	Picture    string    `json:"picture"`
	ExpireDate time.Time `json:"expireDate"`
}

type PhotographerResponse struct {
	ID                uint   `json:"id"`
	Name              string `json:"name"`
	Email             string `json:"email"`
	PhoneNumber       string `json:"phoneNumber"`
	ProfilePictureURL string `json:"profilePictureUrl"`
	IsVerified        bool   `json:"isVerified"`
	ActiveStatus      bool   `json:"activeStatus"`
}

type PhotographerRequest struct {
	PaginationRequest
	Name string `query:"name" default:""`
}

func ToPhotographerResponse(photographer model.Photographer) PhotographerResponse {
	return PhotographerResponse{
		ID:                photographer.UserID,
		Name:              photographer.User.Name,
		Email:             photographer.User.Email,
		PhoneNumber:       photographer.User.PhoneNumber,
		ProfilePictureURL: photographer.User.ProfilePictureURL,
		IsVerified:        photographer.IsVerified,
		ActiveStatus:      photographer.ActiveStatus,
	}
}
