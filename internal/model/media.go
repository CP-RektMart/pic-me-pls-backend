package model

import "gorm.io/gorm"

type Media struct {
	gorm.Model
	GalleryID   uint    `gorm:"not null"`
	Gallery     Gallery `gorm:"foreignKey:GalleryID"`
	PictureURL  string  `gorm:"not null"`
	Description string
}
