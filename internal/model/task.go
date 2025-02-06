package model

import (
	"time"

	"gorm.io/gorm"
)

type Task struct {
	gorm.Model
	QuotationID uint `gorm:"not null"`
	EndTime     time.Time
}
