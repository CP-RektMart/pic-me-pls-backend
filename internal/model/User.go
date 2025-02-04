package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username        string         `gorm:"unique;not null;size:50"`
	Email           string         `gorm:"unique;not null;size:255"`
	PhoneNumber     string         `gorm:"unique;not null;size:10"`
	Password        string         `gorm:"not null;size:100"`
	ProfilePicture  string         `gorm:"size:255"`
	IsBanned        bool           `gorm:"default:false"`

	Photographer    Photographer   `gorm:"foreignKey:UserID"`
	Galleries       []Gallery      `gorm:"foreignKey:UserID"`
	Chats           []Chat         `gorm:"many2many:user_chats;"`
	Messages        []Message      `gorm:"foreignKey:SenderID"`
	Quotations      []Quotation    `gorm:"foreignKey:UserID"`
	Reviews         []Review       `gorm:"foreignKey:UserID"`
}