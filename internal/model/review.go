package model

import (
	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	PackageID   uint      `gorm:"not null"`
	Package     Package   `gorm:"foreignKey:PackageID"`
	CustomerID  uint      `gorm:"not null"`
	Customer    User      `gorm:"foreignKey:CustomerID"`
	Rating      float64   `gorm:"not null"`
	QuotationID uint      `gorm:"not null"`
	Quotation   Quotation `gorm: "foreignKey:QuotationID"`
	IsEdited    bool      `gorm:"not null;default: false"`
	Comment     string
}
