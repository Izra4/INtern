package entity

import "gorm.io/gorm"

type Gedung struct {
	gorm.Model
	Nama      string `gorm:"type:VARCHAR(50);NOT NULL"`
	Alamat    string `gorm:"type:VARCHAR(100);NOT NULL"`
	Kecamatan string `gorm:"type:VARCHAR(50);NOT NULL"`
	Harga     int    `gorm:"type:INT;NOT NULL"`
	Kapasitas string `gorm:"type:VARCHAR(20);NOT NULL"`
	Luas      string `gorm:"type:VARCHAR(15);NOT NULL"`
	Tag       string `gorm:"type:TEXT;NOT NULL"`
	Fasilitas string `gorm:"type:TEXT;NOT NULL"`
	Aturan    string `gorm:"type:TEXT;NOT NULL"`
	Links     []Link
	Bookings  []Booking
	Payments  []Payment
}
