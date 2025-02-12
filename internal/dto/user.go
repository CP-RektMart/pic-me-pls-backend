package dto

type UserUpdateRequest struct {
	Name              string `form:"name"`
	Email             string `form:"email"`
	PhoneNumber       string `form:"phone_number"`
	Facebook          string `form:"facebook"`
	Instagram         string `form:"instagram"`
	Bank              string `form:"bank"`
	AccountNo         string `form:"account_no"`
	BankBranch        string `form:"bank_branch"`
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
