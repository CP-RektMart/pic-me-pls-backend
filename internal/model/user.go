package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRole string

const (
	UserRoleUnknown      UserRole = ""
	UserRoleAdmin        UserRole = "ADMIN"
	UserRolePhotographer UserRole = "PHOTOGRAPHER"
	UserRoleCustomer     UserRole = "CUSTOMER"
)

func (r UserRole) String() string {
	return string(r)
}

func ValidateRole(role string) bool {
	switch UserRole(role) {
	case UserRoleAdmin, UserRolePhotographer, UserRoleCustomer:
		return true
	}
	return false
}
type Provider string

const (
	ProviderUnknown Provider = ""
	ProviderGoogle  Provider = "GOOGLE"
)

func (p Provider) String() string {
	return string(p)
}

func ValidateProvider(provider string) bool {
	switch Provider(provider) {
	case ProviderGoogle:
		return true
	}
	return false
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
	Email             string `gorm:"not null;unique"`
	PhoneNumber       string `gorm:"size:10"`
	ProfilePictureURL string
	Role              UserRole    `gorm:"not null"`
	SenderMessages    []Message   `gorm:"foreignKey:SenderID"`
	ReceiverMessages  []Message   `gorm:"foreignKey:ReceiverID"`
	Quotations        []Quotation `gorm:"foreignKey:CustomerID"`
	Reviews           []Review    `gorm:"foreignKey:CustomerID"`
}

type UserDto struct {
	ID                uint   `json:"id"`
	Name              string `json:"name"`
	PhoneNumber       string `json:"phone_number"`
	ProfilePictureURL string `json:"profile_picture_url"`
	Role              string `json:"role"`
}
