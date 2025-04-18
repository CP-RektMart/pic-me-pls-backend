package dto

import (
	"github.com/CP-RektMart/pic-me-pls-backend/internal/model"
	"github.com/samber/lo"
)

type UserUpdateRequest struct {
	Name              string `json:"name"`
	ProfilePictureURL string `json:"profilePictureUrl"`
	PhoneNumber       string `json:"phoneNumber"`
	Facebook          string `json:"facebook"`
	Instagram         string `json:"instagram"`
	Bank              string `json:"bank"`
	AccountNo         string `json:"accountNo"`
	BankBranch        string `json:"bankBranch"`
}

type UserResponse struct {
	ID                uint           `json:"id"`
	Name              string         `json:"name"`
	Email             string         `json:"email"`
	PhoneNumber       string         `json:"phoneNumber"`
	ProfilePictureURL string         `json:"profilePictureUrl"`
	Role              model.UserRole `json:"role"`
	Facebook          string         `json:"facebook,omitempty"`
	Instagram         string         `json:"instagram,omitempty"`
	Bank              string         `json:"bank,omitempty"`
	AccountNo         string         `json:"accountNo,omitempty"`
	BankBranch        string         `json:"bankBranch,omitempty"`
}

type CustomerResponse struct {
	ID                uint   `json:"id"`
	Name              string `json:"name"`
	Email             string `json:"email"`
	PhoneNumber       string `json:"phoneNumber"`
	ProfilePictureURL string `json:"profilePictureUrl"`
}

type PublicUserResponse struct {
	ID                uint           `json:"id"`
	Name              string         `json:"name"`
	Email             string         `json:"email"`
	PhoneNumber       string         `json:"phoneNumber"`
	ProfilePictureURL string         `json:"profilePictureUrl"`
	Role              model.UserRole `json:"role"`
}

type GetUsersRequest struct {
	PaginationRequest
	Name string `query:"name" default:""`
}

type GetUserByIDRequest struct {
	ID uint `params:"id" validate:"required"`
}

func ToUserResponse(user model.User) UserResponse {
	return UserResponse{
		ID:                user.ID,
		Name:              user.Name,
		Email:             user.Email,
		PhoneNumber:       user.PhoneNumber,
		ProfilePictureURL: user.ProfilePictureURL,
		Role:              user.Role,
		Facebook:          user.Facebook,
		Instagram:         user.Instagram,
		Bank:              user.Bank,
		AccountNo:         user.AccountNo,
		BankBranch:        user.BankBranch,
	}
}

func ToCustomerResponse(user model.User) CustomerResponse {
	return CustomerResponse{
		ID:                user.ID,
		Name:              user.Name,
		Email:             user.Email,
		PhoneNumber:       user.PhoneNumber,
		ProfilePictureURL: user.ProfilePictureURL,
	}
}

func ToPublicUserResponse(user model.User) PublicUserResponse {
	return PublicUserResponse{
		ID:                user.ID,
		Name:              user.Name,
		Email:             user.Email,
		PhoneNumber:       user.PhoneNumber,
		ProfilePictureURL: user.ProfilePictureURL,
		Role:              user.Role,
	}
}

func ToPublicUserResponses(users []model.User) []PublicUserResponse {
	return lo.Map(users, func(u model.User, _ int) PublicUserResponse {
		return ToPublicUserResponse(u)
	})
}
