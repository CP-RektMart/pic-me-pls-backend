package model

import "time"

type Review struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"not null"`
	GalleryID uint      `gorm:"not null"`
	Rating    int       `gorm:"not null;check:rating>=1 AND rating<=5"`
	Comment   string    `gorm:"size:500"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
