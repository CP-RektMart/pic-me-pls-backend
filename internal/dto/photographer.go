package dto

import "time"

type CitizenCardRequest struct {
	CitizenID  string    `json:"citizenId" validate:"required"`
	LaserID    string    `json:"laserId" validate:"required"`
	Picture    string    `json:"picture" validate:"required"`
	ExpireDate time.Time `json:"expireDate" validate:"required"`
}

type CitizenCardResponse struct {
	CitizenID  string    `json:"citizenId"`
	LaserID    string    `json:"laserId"`
	Picture    string    `json:"picture"`
	ExpireDate time.Time `json:"expireDate"`
}
