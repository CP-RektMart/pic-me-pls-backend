package model

import "time"

type Quotation struct {
	ID            uint      `gorm:"primaryKey"`
	UserID        uint      `gorm:"not null"`
	GalleryID     uint      `gorm:"not null"`
	Price         float64   `gorm:"not null"`
	MeetingLoc    string    `gorm:"size:255"`
	StartTime     time.Time `gorm:"not null"`
	EndTime       time.Time `gorm:"not null"`
	DueDate       time.Time `gorm:"not null"`
	TransactionID uint      `gorm:"not null"`
}
