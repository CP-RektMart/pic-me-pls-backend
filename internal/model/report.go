package model

import (
	"time"

	"gorm.io/gorm"
)

type Report struct {
	gorm.Model
	QuotationID uint      `gorm:"not null"`
	Quotation   Quotation `gorm:"foreignKey:QuotationID"`
	CustomerID  uint      `gorm:"not null"`
	Customer    User      `gorm:"foreignKey:CustomerID"`
	Status      string    `gorm:"not null"` // reported, reviewed
	Message     string    `gorm:"not null"`
	DateCreated time.Time `gorm:"not null"`
}
