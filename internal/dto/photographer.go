package dto

import "time"

type CitizenCardRequest struct {
	CitizenID  string    `form:"citizenId" validate:"required"`
	LaserID    string    `form:"laserId" validate:"required"`
	ExpireDate time.Time `form:"expireDate" validate:"required"`
}

type CitizenCardResponse struct {
	CitizenID  string    `json:"citizenId"`
	LaserID    string    `json:"laserId"`
	Picture    string    `json:"picture"`
	ExpireDate time.Time `json:"expireDate"`
}
