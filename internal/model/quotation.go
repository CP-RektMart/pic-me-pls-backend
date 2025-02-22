package model

import (
	"gorm.io/gorm"
)

type QuotationStatus string

func (q QuotationStatus) String() string {
	return string(q)
}

const QuotationPending QuotationStatus = "PENDING"
const QuotationConfirm QuotationStatus = "CONFIRMED"
const QuotationCancelled QuotationStatus = "CANCELLED"
const QuotationPaid QuotationStatus = "PAID"

type Quotation struct {
	gorm.Model
	GalleryID      uint            `gorm:"not null"`
	Gallery        Gallery         `gorm:"foreignKey:GalleryID"`
	CustomerID     uint            `gorm:"not null"`
	Customer       User            `gorm:"foreignKey:CustomerID"`
	PhotographerID uint            `gorm:"not null"`
	Photographer   Photographer    `gorm:"foreignKey:PhotographerID"`
	Status         QuotationStatus `gorm:"not null"` // pending, confirm, cancelled, paid
	Price          float64         `gorm:"not null"`
	Description    string          `gorm:"not null"`
}
