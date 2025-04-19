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

type ListUnverifiedPhotographerResponse struct {
	ID                uint                `json:"id"`
	Name              string              `json:"name"`
	Email             string              `json:"email"`
	PhoneNumber       string              `json:"phoneNumber"`
	ProfilePictureURL string              `json:"profilePictureUrl"`
	IsVerified        bool                `json:"isVerified"`
	ActiveStatus      bool                `json:"activeStatus"`
	CitizenCard       CitizenCardResponse `json:"citizenCard"`
}

func ToListUnverifiedPhotographerResponse(p model.Photographer) ListUnverifiedPhotographerResponse {
	return ListUnverifiedPhotographerResponse{
		ID:                p.UserID,
		Name:              p.User.Name,
		Email:             p.User.Email,
		PhoneNumber:       p.User.PhoneNumber,
		ProfilePictureURL: p.User.ProfilePictureURL,
		IsVerified:        p.IsVerified,
		ActiveStatus:      p.ActiveStatus,
		CitizenCard:       ToCitizenCardResponse(&p.CitizenCard),
	}
}

func ToListUnverifiedPhotographersResponse(ps []model.Photographer) []ListUnverifiedPhotographerResponse {
	return lo.Map(ps, func(p model.Photographer, _ int) ListUnverifiedPhotographerResponse {
		return ToListUnverifiedPhotographerResponse(p)
	})
}
