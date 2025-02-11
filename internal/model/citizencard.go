package model

import (
	"time"

	"gorm.io/gorm"
)

type CitizenCard struct {
	gorm.Model
	CitizenID  string    `gorm:"size:255;not null"`
	LaserID    string    `gorm:"size:255;not null"`
	Picture    string    `gorm:"size:255;not null"`
	ExpireDate time.Time `gorm:"not null"`
}
