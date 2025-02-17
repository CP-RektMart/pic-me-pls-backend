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
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	IsVerified   bool   `json:"is_verified"`
	ActiveStatus bool   `json:"active_status"`
}

func ToPhotographerResponse(photographer model.Photographer) PhotographerResponse {
	return PhotographerResponse{
		ID:           photographer.ID,
		Name:         photographer.User.Name,
		IsVerified:   photographer.IsVerified,
		ActiveStatus: photographer.ActiveStatus,
	}
}
