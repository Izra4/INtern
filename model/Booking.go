package model

import "time"

type Booking struct {
	Nama      string    `json:"nama" binding:"required"`
	Tanggal   time.Time `json:"tanggal" binding:"required"`
	Keperluan string    `json:"keperluan" binding:"required"`
	Nomor     string    `json:"nomor" binding:"required"`
	Alamat    string    `json:"alamat" binding:"required"`
	Fasilitas string    `json:"fasilitas" binding:"required"`
	UserID    uint
	GedungID  uint
}
