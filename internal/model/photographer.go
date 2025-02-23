package model

import "gorm.io/gorm"

type Photographer struct {
	gorm.Model
	UserID uint `gorm:"uniqueIndex;not null"`
	User   User `gorm:"foreignKey:UserID"`

	IsVerified   bool      `gorm:"not null;default:false"`
	ActiveStatus bool      `gorm:"not null;default:false"`
	Packages     []Package `gorm:"foreignKey:PhotographerID"`

	CitizenCardID *uint       `gorm:"uniqueIndex"`
	CitizenCard   CitizenCard `gorm:"foreignKey:CitizenCardID"`
}
