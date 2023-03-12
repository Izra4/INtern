package entity

import "gorm.io/gorm"

type Gedung struct {
	gorm.Model
	Nama      string `gorm:"type:VARCHAR(50);NOT NULL" json:"nama"`
	Alamat    string `gorm:"type:VARCHAR(100);NOT NULL" json:"alamat"`
	Kecamatan string `gorm:"type:VARCHAR(50);NOT NULL" json:"kecamatan"`
	Harga     int    `gorm:"type:INT;NOT NULL" json:"harga"`
	Kapasitas string `gorm:"type:VARCHAR(20);NOT NULL" json:"kapasitas"`
	Luas      string `gorm:"type:VARCHAR(15);NOT NULL" json:"luas"`
	Tag       string `gorm:"type:TEXT;NOT NULL" json:"tag"`
	Fasilitas string `gorm:"type:TEXT;NOT NULL" json:"fasilitas"`
	Aturan    string `gorm:"type:TEXT;NOT NULL" json:"aturan"`
	Links     []Link
	Bookings  []Booking
	Payments  []Payment
}
