package dto

type UserRequest struct {
	ID                uint   `json:"id" validate:"required"`
	Name              string `json:"name" validate:"required"`
	Email             string `json:"email" validate:"required"`
	PhoneNumber       string `json:"phone_number" validate:"required"`
	ProfilePictureURL string `json:"profile_picture_url"`
	Role              string `json:"role" validate:"required"`
	Facebook          string `json:"facebook"`
	Instagram         string `json:"instagram"`
	Bank              string `json:"bank"`
	AccountNo         string `json:"account_no"`
	BankBranch        string `json:"bank_branch"`
}

type UserResponse struct {
	ID                uint   `json:"id"`
	Name              string `json:"name"`
	Email             string `json:"email"`
	PhoneNumber       string `json:"phone_number"`
	ProfilePictureURL string `json:"profile_picture_url"`
	Role              string `json:"role"`
	Facebook          string `json:"facebook,omitempty"`
	Instagram         string `json:"instagram,omitempty"`
	Bank              string `json:"bank,omitempty"`
	AccountNo         string `json:"account_no,omitempty"`
	BankBranch        string `json:"bank_branch,omitempty"`
}
