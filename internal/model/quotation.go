package model

import (
	"gorm.io/gorm"
)

type Quotation struct {
	gorm.Model
	GalleryID      uint         `gorm:"not null"`
	Gallery        Gallery      `gorm:"foreignKey:GalleryID"`
	CustomerID     uint         `gorm:"not null"`
	Customer       User         `gorm:"foreignKey:CustomerID"`
	PhotographerID uint         `gorm:"not null"`
	Photographer   Photographer `gorm:"foreignKey:PhotographerID"`
	Status         string       `gorm:"not null"` // pending, confirm, cancelled, paid
	Price          float64      `gorm:"not null"`
	Description    string       `gorm:"not null"`
}
