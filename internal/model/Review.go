package model

import (
	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	UserID    uint   `gorm:"not null"`
	GalleryID uint   `gorm:"not null"`
	Rating    int    `gorm:"not null;check:rating>=1 AND rating<=5"`
	Comment   string `gorm:"size:500"`
}
