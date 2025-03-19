package model

import "gorm.io/gorm"

type Preview struct {
	gorm.Model
	Link        string    `gorm:"not null"`
	QuotationID uint      `gorm:"not null"`
	Quotation   Quotation `gorm:"foreignKey:QuotationID"`
}
