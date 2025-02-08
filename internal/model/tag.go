package model

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	GalleryID uint    `gorm:"not null"`
	Gallery   Gallery `gorm:"foreignKey:GalleryID"`
	Name      string  `gorm:"size:100;not null"`
}
