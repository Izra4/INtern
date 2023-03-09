package entity

import (
	"gorm.io/gorm"
)

type Booking struct {
	gorm.Model
	Nama      string `gorm:"type:VARCHAR(50);NOT NULL"`
	Tanggal   string `gorm:"type:time;NOT NULL"`
	Keperluan string `gorm:"type:VARCHAR(30);NOT NULL'"`
	Nomer     string `gorm:"type:VARCHAR(15);NOT NULL"`
	Alamat    string `gorm:"type:VARCHAR(200);NOT NULL"`
	UserID    uint
	GedungID  uint
}
