package dto

import (
	"time"

	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/samber/lo"
)

type CitizenCardResponse struct {
	ID         uint      `json:"id"`
	CitizenID  string    `json:"citizenId"`
	LaserID    string    `json:"laserId"`
	Picture    string    `json:"picture"`
	ExpireDate time.Time `json:"expireDate"`
}

func ToCitizenCardResponse(c *model.CitizenCard) CitizenCardResponse {
	return CitizenCardResponse{
		ID:         c.ID,
		CitizenID:  c.CitizenID,
		LaserID:    c.LaserID,
		Picture:    c.Picture,
		ExpireDate: c.ExpireDate,
	}
}

type ListUnverifiedCitizenCardResponse struct {
	ID                uint                `json:"id"`
	Name              string              `json:"name"`
	Email             string              `json:"email"`
	PhoneNumber       string              `json:"phoneNumber"`
	ProfilePictureURL string              `json:"profilePictureUrl"`
	IsVerified        bool                `json:"isVerified"`
	ActiveStatus      bool                `json:"activeStatus"`
	IsBanned          bool                `json:"isBanned"`
	CitizenCard       CitizenCardResponse `json:"citizenCard"`
}

func ToListUnverifiedCitizenCardResponse(c model.CitizenCard) ListUnverifiedCitizenCardResponse {
	return ListUnverifiedCitizenCardResponse{
		ID:                c.Photographer.UserID,
		Name:              c.Photographer.User.Name,
		Email:             c.Photographer.User.Email,
		PhoneNumber:       c.Photographer.User.PhoneNumber,
		ProfilePictureURL: c.Photographer.User.ProfilePictureURL,
		IsVerified:        c.Photographer.IsVerified,
		ActiveStatus:      c.Photographer.ActiveStatus,
		IsBanned:          c.Photographer.IsBanned,
		CitizenCard:       ToCitizenCardResponse(&c),
	}
}

func ToListUnverifiedCitizenCardResponses(citizenCards []model.CitizenCard) []ListUnverifiedCitizenCardResponse {
	return lo.Map(citizenCards, func(c model.CitizenCard, _ int) ListUnverifiedCitizenCardResponse {
		return ToListUnverifiedCitizenCardResponse(c)
	})
}
