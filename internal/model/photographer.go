package model

import "gorm.io/gorm"

type Photographer struct {
	gorm.Model
	UserID           uint   `gorm:"not null;unique"`
	User             User   `gorm:"foreignKey:UserID"`
	SSN              string `gorm:"size:10;not null"`
	IsVerified       bool   `gorm:"not null;default:false"`
	ActiveStatus     bool   `gorm:"not null;default:false"`
	IDCardPictureURL string
}
