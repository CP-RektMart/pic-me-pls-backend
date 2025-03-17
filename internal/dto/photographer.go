package dto

import (
	"time"

	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
)

type VerifyCitizenCardRequest struct {
	CitizenID  string    `json:"citizenId" validate:"required"`
	ImageURL   string    `json:"imageUrl" validate:"required"`
	LaserID    string    `json:"laserId" validate:"required"`
	ExpireDate time.Time `json:"expireDate" validate:"required"`
}

type ReVerifyCitizenCardRequest struct {
	CitizenID  string    `json:"citizenId"`
	ImageURL   string    `json:"imageUrl"`
	LaserID    string    `json:"laserId"`
	ExpireDate time.Time `json:"expireDate"`
}

type CitizenCardResponse struct {
	CitizenID  string    `json:"citizenId"`
	LaserID    string    `json:"laserId"`
	Picture    string    `json:"picture"`
	ExpireDate time.Time `json:"expireDate"`
}

type PhotographerResponse struct {
	ID                uint                   `json:"id"`
	Name              string                 `json:"name"`
	Email             string                 `json:"email"`
	PhoneNumber       string                 `json:"phoneNumber"`
	ProfilePictureURL string                 `json:"profilePictureUrl"`
	IsVerified        bool                   `json:"isVerified"`
	ActiveStatus      bool                   `json:"activeStatus"`
	Packages          []SmallPackageResponse `json:"packages"`
}

type PhotographerRequest struct {
	PaginationRequest
	Name string `query:"name" default:""`
}

type GetPhotographerByIDRequest struct {
	ID uint `params:"id" validate:"required"`
}

func ToPhotographerResponse(photographer model.Photographer) PhotographerResponse {
	return PhotographerResponse{
		ID:                photographer.ID,
		Name:              photographer.User.Name,
		Email:             photographer.User.Email,
		PhoneNumber:       photographer.User.PhoneNumber,
		ProfilePictureURL: photographer.User.ProfilePictureURL,
		IsVerified:        photographer.IsVerified,
		ActiveStatus:      photographer.ActiveStatus,
		Packages:          ToSmallPackageResponses(photographer.Packages),
	}
}
