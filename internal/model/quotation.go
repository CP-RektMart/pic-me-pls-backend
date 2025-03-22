package model

import (
	"time"

	"gorm.io/gorm"
)

type QuotationStatus string

func (q QuotationStatus) String() string {
	return string(q)
}

func (q QuotationStatus) IsValid() bool {
	switch q {
	case QuotationPending, QuotationConfirm, QuotationCancelled, QuotationPaid:
		return true
	default:
		return false
	}
}

const QuotationPending QuotationStatus = "PENDING"
const QuotationConfirm QuotationStatus = "CONFIRMED"
const QuotationCancelled QuotationStatus = "CANCELLED"
const QuotationPaid QuotationStatus = "PAID"
const QuotationAccepted QuotationStatus = "ACCEPTED"

type Quotation struct {
	gorm.Model
	PackageID      uint            `gorm:"not null"`
	Package        Package         `gorm:"foreignKey:PackageID"`
	CustomerID     uint            `gorm:"not null"`
	Customer       User            `gorm:"foreignKey:CustomerID"`
	PhotographerID uint            `gorm:"not null"`
	Photographer   Photographer    `gorm:"foreignKey:PhotographerID"`
	Status         QuotationStatus `gorm:"not null"` // pending, confirm, cancelled, paid
	Price          float64         `gorm:"not null"`
	Description    string          `gorm:"not null"`
	FromDate       time.Time       `gorm:"not null"`
	ToDate         time.Time       `gorm:"not null"`
}
