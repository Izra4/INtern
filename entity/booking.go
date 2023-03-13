package entity

import (
	"gorm.io/gorm"
	"time"
)

type Booking struct {
	gorm.Model
	Nama      string    `gorm:"type:VARCHAR(50);NOT NULL"`
	Tanggal   time.Time `gorm:"type:time;NOT NULL" type:"date"`
	Keperluan string    `gorm:"type:VARCHAR(30);NOT NULL'"`
	Nomer     string    `gorm:"type:VARCHAR(15);NOT NULL"`
	Alamat    string    `gorm:"type:VARCHAR(200);NOT NULL"`
	Fasilitas string    `gorm:"type:VARCHAR(200);NOT NULL"`
	UserID    uint
	GedungID  uint
}
