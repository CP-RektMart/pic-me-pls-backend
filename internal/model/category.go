package model

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name        string    `gorm:"not null"`
	Description string    `gorm:"not null"`
	Galleries   []Gallery `gorm:"many2many:Galleries_Categories"`
}
