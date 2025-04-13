package model

import (
	"time"

	"gorm.io/gorm"
)

type Report struct {
	gorm.Model
	QuotationID    uint      `gorm:"not null"`
	Quotation      Quotation `gorm:"foreignKey:QuotationID"`
	CustomerID     uint      `gorm:"not null"`
	Customer       User      `gorm:"foreignKey:CustomerID"`
	PhotographerID uint      `gorm:"not null"`
	Photographer   User      `gorm:"foreignKey:PhotographerID"`
	ReporterID     uint      `gorm:"not null"`
	Reporter       User      `gorm:"foreignKey:ReporterID"`
	ReporterRole   string    `gorm:"not null"` // customer, photographer
	Status         string    `gorm:"not null"` // reported, reviewed
	Message        string    `gorm:"not null"`
	DateCreated    time.Time `gorm:"not null"`
}
