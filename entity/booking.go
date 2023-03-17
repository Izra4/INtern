package entity

import (
	"gorm.io/gorm"
	"time"
)

type Booking struct {
	gorm.Model
	Nama      string    `gorm:"type:VARCHAR(50);NOT NULL" json:"nama"`
	Tanggal   time.Time `gorm:"type:time;NOT NULL" json:"tanggal"`
	Keperluan string    `gorm:"type:VARCHAR(30);NOT NULL'" json:"keperluan"`
	Nomer     string    `gorm:"type:VARCHAR(15);NOT NULL" json:"nomer"`
	Alamat    string    `gorm:"type:VARCHAR(200);NOT NULL" json:"alamat"`
	Fasilitas string    `gorm:"type:VARCHAR(200);NOT NULL" json:"fasilitas"`
	UserID    uint
	GedungID  uint
	Payment   Payment
}

type BookingPost struct {
	Nama      string `json:"nama" binding:"required"`
	Tanggal   string `json:"tanggal" binding:"required"`
	Keperluan string `json:"keperluan" binding:"required"`
	Nomor     string `json:"nomor" binding:"required"`
	Alamat    string `json:"alamat" binding:"required"`
	Fasilitas string `json:"fasilitas" binding:"required"`
	UserID    uint
	GedungID  uint
}
