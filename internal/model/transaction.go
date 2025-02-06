package model

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	Amount float64 `gorm:"not null"`
}
