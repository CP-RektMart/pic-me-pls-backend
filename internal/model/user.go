package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRole string

const (
	UserRoleUnknown UserRole = ""
	UserRoleAdmin   UserRole = "ADMIN"
	UserRoleUser    UserRole = "USER"
)

func (r UserRole) String() string {
	return string(r)
}

type Provider string

const (
	ProviderUnknown Provider = ""
	ProviderGoogle  Provider = "GOOGLE"
)

func (p Provider) String() string {
	return string(p)
}

type Token struct {
	AccessToken  string
	RefreshToken string
	Exp          int64
}

type CachedTokens struct {
	AccessUID  uuid.UUID
	RefreshUID uuid.UUID
}

type User struct {
	gorm.Model
	Name              string `gorm:"not null"`
	PhoneNumber       string `gorm:"size:10"`
	ProfilePictureURL string
	Role              UserRole    `gorm:"not null"`
	SenderMessages    []Message   `gorm:"foreignKey:SenderID"`
	ReceiverMessages  []Message   `gorm:"foreignKey:ReceiverID"`
	Quotations        []Quotation `gorm:"foreignKey:CustomerID"`
	Reviews           []Review    `gorm:"foreignKey:CustomerID"`
}
