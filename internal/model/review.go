package model

import (
	"time"

	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	ID         uint    `gorm:"primaryKey"`
	GalleryID  uint    `gorm:"not null"`
	Gallery    Gallery `gorm:"foreignKey:GalleryID"`
	CustomerID uint    `gorm:"not null"`
	Customer   User    `gorm:"foreignKey:CustomerID"`
	Rating     float64 `gorm:"not null"`
	Comment    string
	CreatedAt  time.Time `gorm:"not null"`
}
