package model

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name        string    `gorm:"not null"`
	Description string    `gorm:"not null"`
	Packages    []Package `gorm:"many2many:Packages_Categories"`
}
