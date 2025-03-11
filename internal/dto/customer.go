package dto

import "github.com/CP-RektMart/pic-me-pls-backend/internal/model"

type GetCustomerPublicRequest struct {
	ID uint `params:"id" validate:"required"`
}

type CustomerPublicResponse struct {
	ID                uint   `json:"id"`
	Name              string `json:"name"`
	ProfilePictureURL string `json:"profilePictureUrl"`
}

func ToCustomerPublicResponse(user model.User) CustomerPublicResponse {
	return CustomerPublicResponse{
		ID:                user.ID,
		Name:              user.Name,
		ProfilePictureURL: user.ProfilePictureURL,
	}
}
