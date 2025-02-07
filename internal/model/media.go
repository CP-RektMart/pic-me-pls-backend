package model

import "gorm.io/gorm"

type Media struct {
	gorm.Model
	ID         uint    `gorm:"primaryKey"`
	GalleryID  uint    `gorm:"not null"`
	Gallery    Gallery `gorm:"foreignKey:GalleryID"`
	PictureURL string  `gorm:"not null"`
}
