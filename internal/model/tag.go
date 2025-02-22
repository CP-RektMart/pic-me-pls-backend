package model

import "gorm.io/gorm"

type Tag struct {
	gorm.Model
	PackageID uint    `gorm:"not null"`
	Package   Package `gorm:"foreignKey:PackageID"`
	Name      string  `gorm:"size:100;not null"`
}
