package model

import "gorm.io/gorm"

type Gallery struct {
	gorm.Model
	PhotographerID uint         `gorm:"not null"`
	Photographer   Photographer `gorm:"foreignKey:PhotographerID"`
	Name           string       `gorm:"not null"`
	Description    string
	Price          float64     `gorm:"not null"`
	Tags           []Tag       `gorm:"foreignKey:GalleryID"`
	Media          []Media     `gorm:"foreignKey:GalleryID"`
	Reviews        []Review    `gorm:"foreignKey:GalleryID"`
	Categories     []Category  `gorm:"many2many:Galleries_Categories"`
	Quotations     []Quotation `gorm:"foreignKey:GalleryID"`
}
