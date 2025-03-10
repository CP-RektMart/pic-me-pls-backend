package model

import "gorm.io/gorm"

type Package struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Description string
	Price       float64 `gorm:"not null"`

	CategoryID uint
	Category   Category

	PhotographerID uint         `gorm:"not null"`
	Photographer   Photographer `gorm:"foreignKey:PhotographerID"`

	Tags       []Tag
	Media      []Media
	Reviews    []Review
	Quotations []Quotation
}
