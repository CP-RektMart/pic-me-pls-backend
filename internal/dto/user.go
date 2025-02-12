package dto

type UserUpdateRequest struct {
	Name        string
	PhoneNumber string
	Facebook    string
	Instagram   string
	Bank        string
	AccountNo   string
	BankBranch  string
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
