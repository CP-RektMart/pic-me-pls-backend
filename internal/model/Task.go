package model

import "time"

type Task struct {
	ID          uint      `gorm:"primaryKey"`
	QuotationID uint      `gorm:"not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	EndTime     time.Time
}
