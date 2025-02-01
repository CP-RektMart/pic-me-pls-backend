package model

import "time"

type Media struct {
	ID        uint      `gorm:"primaryKey"`
	GalleryID uint      `gorm:"not null"`
	URL       string    `gorm:"size:255;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	Gallery   Gallery   `gorm:"foreignKey:GalleryID"`
}
