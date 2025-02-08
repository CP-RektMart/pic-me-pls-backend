package model

import (
	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	GalleryID  uint    `gorm:"not null"`
	Gallery    Gallery `gorm:"foreignKey:GalleryID"`
	CustomerID uint    `gorm:"not null"`
	Customer   User    `gorm:"foreignKey:CustomerID"`
	Rating     float64 `gorm:"not null"`
	Comment    string
}
