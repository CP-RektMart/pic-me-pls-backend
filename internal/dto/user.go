package dto

type BaseUserDTO struct {
	ID                uint   `json:"id"`
	Name              string `json:"name"`
	Email             string `json:"email"`
	PhoneNumber       string `json:"phone_number"`
	ProfilePictureURL string `json:"profile_picture_url"`
	Role              string `json:"role"`
}

type UpdateUserRequest struct {
	BaseUserDTO
	Facebook   string `json:"facebook,omitempty"`
	Instagram  string `json:"instagram,omitempty"`
	Bank       string `json:"bank,omitempty"`
	AccountNo  string `json:"account_no,omitempty"`
	BankBranch string `json:"bank_branch,omitempty"`
}

type UpdateUserResponse struct {
	BaseUserDTO
}
