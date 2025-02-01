package model

import "time"

type Transaction struct {
	ID        uint      `gorm:"primaryKey"`
	Amount    float64   `gorm:"not null"`
	Timestamp time.Time `gorm:"autoCreateTime"`
}
