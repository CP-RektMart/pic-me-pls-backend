package model

import "time"

type Gallery struct {
	ID           uint         `gorm:"primaryKey"`
	UserID       uint         `gorm:"not null"`
	TypeID       uint         `gorm:"not null"`
	Name         string       `gorm:"size:100;not null"`
	Description  string       `gorm:"size:500"`
	Equipment    string       `gorm:"size:255"`
	Price        float64      `gorm:"not null"`
	CreatedAt    time.Time    `gorm:"autoCreateTime"`
	Photographer Photographer `gorm:"foreignKey:UserID"`
}
