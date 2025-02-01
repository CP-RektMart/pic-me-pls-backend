package model

type Photographer struct {
	UserID        uint   `gorm:"primaryKey"`
	SSN           string `gorm:"unique;not null;size:13"`
	IDCardPicture string `gorm:"size:255"`
	IsVerified    bool   `gorm:"default:false"`
	ActiveStatus  bool   `gorm:"default:true"`
	User          User   `gorm:"foreignKey:UserID"`
}
