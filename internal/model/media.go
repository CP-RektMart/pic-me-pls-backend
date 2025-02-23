package model

import "gorm.io/gorm"

type Media struct {
	gorm.Model
	PackageID   uint    `gorm:"not null"`
	Package     Package `gorm:"foreignKey:PackageID"`
	PictureURL  string  `gorm:"not null"`
	Description string
}
