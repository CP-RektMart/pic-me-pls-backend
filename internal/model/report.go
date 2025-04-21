package model

import (
	"gorm.io/gorm"
)

type Report struct {
	gorm.Model
	QuotationID uint         `gorm:"not null"`
	Quotation   Quotation    `gorm:"foreignKey:QuotationID"`
	ReporterID  uint         `gorm:"not null"`
	Reporter    User         `gorm:"foreignKey:ReporterID"`
	Status      ReportStatus `gorm:"not null"`
	Message     string       `gorm:"not null"`
	Title       string       `gorm:"not null"`
}

type ReportStatus string

const (
	ReportStatusAccepted    ReportStatus = "ACCEPTED"
	ReportStatusDeclined ReportStatus = "DECLINED"
)
