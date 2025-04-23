package model

import (
	"time"

	"gorm.io/gorm"
)

type Photographer struct {
	UserID    uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	User User `gorm:"foreignKey:UserID;references:ID"`

	IsVerified   bool      `gorm:"not null;default:false"`
	ActiveStatus bool      `gorm:"not null;default:false"`
	IsBanned     bool      `gorm:"not null;default:false"`
	Packages     []Package `gorm:"foreignKey:PhotographerID"`

	CitizenCard CitizenCard
}
