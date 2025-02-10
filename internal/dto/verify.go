package dto

import "time"

type VerifyCardRequest struct {
	CitizenID  string    `json:"citizen_id" validate:"required"`
	LaserID    string    `json:"laser_id" validate:"required"`
	Picture    string    `json:"picture" validate:"required"`
	ExpireDate time.Time `json:"expire_date" validate:"required"`
}
