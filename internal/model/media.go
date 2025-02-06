package model

import (
	"gorm.io/gorm"
)

type Media struct {
	gorm.Model
	GalleryID uint    `gorm:"not null"`
	URL       string  `gorm:"size:255;not null"`
	Gallery   Gallery `gorm:"foreignKey:GalleryID"`
}
