package dto

import (
	"time"
)

// Form data
type CitizenCardRequest struct {
	CitizenID  string    `validate:"required"`
	LaserID    string    `validate:"required"`
	ExpireDate time.Time `validate:"required"`
}

type CitizenCardResponse struct {
	CitizenID  string    `json:"citizenId"`
	LaserID    string    `json:"laserId"`
	Picture    string    `json:"picture"`
	ExpireDate time.Time `json:"expireDate"`
}
