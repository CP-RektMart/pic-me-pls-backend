package model

import "time"

type Chat struct {
	ID        uint      `gorm:"primaryKey"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	Messages  []Message `gorm:"foreignKey:ChatID"`
}
