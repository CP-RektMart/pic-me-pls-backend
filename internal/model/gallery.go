package model

import "gorm.io/gorm"

type Gallery struct {
	gorm.Model
	PhotographerID uint         `gorm:"not null"`
	Photographer   Photographer `gorm:"foreignKey:PhotographerID"`
	Name           string       `gorm:"not null"`
	Description    string
	Price          float64   `gorm:"not null"`
	Gallery        []Gallery `gorm:"foreignKey:GalleryID"`
}
