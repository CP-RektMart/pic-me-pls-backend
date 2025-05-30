package dto

import (
	"time"

	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/samber/lo"
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

type PhotographerResponse struct {
	ID                uint                   `json:"id"`
	Name              string                 `json:"name"`
	Email             string                 `json:"email"`
	PhoneNumber       string                 `json:"phoneNumber"`
	ProfilePictureURL string                 `json:"profilePictureUrl"`
	IsVerified        bool                   `json:"isVerified"`
	ActiveStatus      bool                   `json:"activeStatus"`
	IsBanned          bool                   `json:"isBanned"`
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
		ID:                photographer.UserID,
		Name:              photographer.User.Name,
		Email:             photographer.User.Email,
		PhoneNumber:       photographer.User.PhoneNumber,
		ProfilePictureURL: photographer.User.ProfilePictureURL,
		IsVerified:        photographer.IsVerified,
		ActiveStatus:      photographer.ActiveStatus,
		IsBanned:          photographer.IsBanned,
		Packages:          ToSmallPackageResponses(photographer.Packages),
	}
}

type ListPhotographerResponse struct {
	ID                uint   `json:"id"`
	Name              string `json:"name"`
	Email             string `json:"email"`
	PhoneNumber       string `json:"phoneNumber"`
	ProfilePictureURL string `json:"profilePictureUrl"`
	IsVerified        bool   `json:"isVerified"`
	IsBanned          bool   `json:"isBanned"`
	ActiveStatus      bool   `json:"activeStatus"`
}

func ToListPhotographerResponse(p model.Photographer) ListPhotographerResponse {
	return ListPhotographerResponse{
		ID:                p.UserID,
		Name:              p.User.Name,
		Email:             p.User.Email,
		PhoneNumber:       p.User.PhoneNumber,
		ProfilePictureURL: p.User.ProfilePictureURL,
		IsVerified:        p.IsVerified,
		IsBanned:          p.IsBanned,
		ActiveStatus:      p.ActiveStatus,
	}
}

func ToListPhotographersResponse(ps []model.Photographer) []ListPhotographerResponse {
	return lo.Map(ps, func(p model.Photographer, _ int) ListPhotographerResponse {
		return ToListPhotographerResponse(p)
	})
}

type DetailPhotographerResponse struct {
	ID                uint                  `json:"id"`
	Name              string                `json:"name"`
	Email             string                `json:"email"`
	PhoneNumber       string                `json:"phoneNumber"`
	ProfilePictureURL string                `json:"profilePictureUrl"`
	IsVerified        bool                  `json:"isVerified"`
	ActiveStatus      bool                  `json:"activeStatus"`
	IsBanned          bool                  `json:"isBanned"`
	Packages          []ListPackageResponse `json:"packages"`
}

func ToDetailPhotographerResponse(p model.Photographer) DetailPhotographerResponse {
	return DetailPhotographerResponse{
		ID:                p.UserID,
		Name:              p.User.Name,
		Email:             p.User.Email,
		PhoneNumber:       p.User.PhoneNumber,
		ProfilePictureURL: p.User.ProfilePictureURL,
		IsVerified:        p.IsVerified,
		ActiveStatus:      p.ActiveStatus,
		IsBanned:          p.IsBanned,
		Packages:          ToListPackagesResponse(p.Packages),
	}
}

type PhotographerMeResponse struct {
	ID                uint   `json:"id"`
	Name              string `json:"name"`
	Email             string `json:"email"`
	PhoneNumber       string `json:"phoneNumber"`
	ProfilePictureURL string `json:"profilePictureUrl"`
	IsVerified        bool   `json:"isVerified"`
	ActiveStatus      bool   `json:"activeStatus"`
	IsBanned          bool   `json:"isBaned"`
}

func ToPhotographerMeResponse(p model.Photographer) PhotographerMeResponse {
	return PhotographerMeResponse{
		ID:                p.UserID,
		Name:              p.User.Name,
		Email:             p.User.Email,
		PhoneNumber:       p.User.PhoneNumber,
		ProfilePictureURL: p.User.ProfilePictureURL,
		IsVerified:        p.IsVerified,
		ActiveStatus:      p.ActiveStatus,
		IsBanned:          p.IsBanned,
	}
}
