package model

import "time"

type User struct {
	ID             uint      `gorm:"primaryKey"`
	Username       string    `gorm:"unique;not null;size:50"`
	Email          string    `gorm:"unique;not null;size:255"`
	PhoneNumber    string    `gorm:"unique;not null;size:10"`
	Password       string    `gorm:"not null;size:100"`
	ProfilePicture string    `gorm:"size:255"`
	IsBanned       bool      `gorm:"default:false"`
	CreatedAt      time.Time `gorm:"autoCreateTime"`
}
