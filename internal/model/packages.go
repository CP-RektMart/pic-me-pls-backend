package model

import "gorm.io/gorm"

type Package struct {
	gorm.Model
	PhotographerID uint         `gorm:"not null"`
	Photographer   Photographer `gorm:"foreignKey:PhotographerID"`
	Name           string       `gorm:"not null"`
	Description    string
	Price          float64     `gorm:"not null"`
	Tags           []Tag       `gorm:"foreignKey:PackageID"`
	Media          []Media     `gorm:"foreignKey:PackageID"`
	Reviews        []Review    `gorm:"foreignKey:PackageID"`
	Categories     []Category  `gorm:"many2many:Packages_Categories"`
	Quotations     []Quotation `gorm:"foreignKey:PackageID"`
}
