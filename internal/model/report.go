package model

import (
	"gorm.io/gorm"
)

type Report struct {
	gorm.Model
	QuotationID  uint      `gorm:"not null"`
	Quotation    Quotation `gorm:"foreignKey:QuotationID"`
	ReporterID   uint      `gorm:"not null"`
	Reporter     User      `gorm:"foreignKey:ReporterID"`
	ReporterRole string    `gorm:"not null"` // customer, admin
	Status       string    `gorm:"not null"` // reported, reviewed
	Message      string    `gorm:"not null"`
}
