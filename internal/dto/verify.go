package dto

import "time"

type VerifyCardRequest struct {
	CitizenID  string    `json:"citizen_id"`
	LaserID    string    `json:"laser_id"`
	Picture    string    `json:"picture"`
	ExpireDate time.Time `json:"expire_date"`
}

type VerifyCardResponse struct {
	Result any    `json:"result" doc:"Result"`
	Error  string `json:"error,omitempty" doc:"Error message"`
}
