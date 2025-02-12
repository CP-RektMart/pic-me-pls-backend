package dto

import "github.com/CP-RektMart/pic-me-pls-backend/internal/model"

type UserUpdateRequest struct {
	Name        string `form:"name"`
	Email       string `form:"email"`
	PhoneNumber string `form:"phoneNumber"`
	Facebook    string `form:"facebook"`
	Instagram   string `form:"instagram"`
	Bank        string `form:"bank"`
	AccountNo   string `form:"accountNo"`
	BankBranch  string `form:"bankBranch"`
}

type UserResponse struct {
	ID                uint   `json:"id"`
	Name              string `json:"name"`
	Email             string `json:"email"`
	PhoneNumber       string `json:"phoneNumber"`
	ProfilePictureURL string `json:"profilePictureUrl"`
	Role              string `json:"role"`
	Facebook          string `json:"facebook,omitempty"`
	Instagram         string `json:"instagram,omitempty"`
	Bank              string `json:"bank,omitempty"`
	AccountNo         string `json:"accountNo,omitempty"`
	BankBranch        string `json:"bankBranch,omitempty"`
}

func ToUserResponse(user model.User) UserResponse {
	return UserResponse{
		ID:                user.ID,
		Name:              user.Name,
		Email:             user.Email,
		PhoneNumber:       user.PhoneNumber,
		ProfilePictureURL: user.ProfilePictureURL,
		Role:              user.Role.String(),
		Facebook:          user.Facebook,
		Instagram:         user.Instagram,
		Bank:              user.Bank,
		AccountNo:         user.AccountNo,
		BankBranch:        user.BankBranch,
	}
}
