package model

import (
	"gorm.io/gorm"
)

type Gallery struct {
	gorm.Model
	UserID       uint         `gorm:"not null"`
	TypeID       uint         `gorm:"not null"`
	Name         string       `gorm:"size:100;not null"`
	Description  string       `gorm:"size:500"`
	Equipment    string       `gorm:"size:255"`
	Price        float64      `gorm:"not null"`
	Photographer Photographer `gorm:"foreignKey:UserID"`
}
