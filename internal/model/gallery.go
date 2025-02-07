package model

import "gorm.io/gorm"

type Gallery struct {
	gorm.Model
	ID             uint         `gorm:"primaryKey"`
	PhotographerID uint         `gorm:"not null"`
	Photographer   Photographer `gorm:"foreignKey:PhotographerID"`
	Name           string       `gorm:"not null"`
	Description    string
	Price          float64 `gorm:"not null"`
}
