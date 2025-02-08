package model

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Description string `gorm:"not null"`
	Tags        []Tag  `gorm:"foreignKey:TagID"`
}
