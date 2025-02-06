package model

import (
	"time"

	"gorm.io/gorm"
)

type Quotation struct {
	gorm.Model
	UserID        uint      `gorm:"not null"`
	GalleryID     uint      `gorm:"not null"`
	Price         float64   `gorm:"not null"`
	MeetingLoc    string    `gorm:"size:255"`
	StartTime     time.Time `gorm:"not null"`
	EndTime       time.Time `gorm:"not null"`
	DueDate       time.Time `gorm:"not null"`
	TransactionID uint      `gorm:"not null"`
}
