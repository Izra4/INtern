package entity

import "gorm.io/gorm"

type Gedung struct {
	gorm.Model
	Nama      string `gorm:"type:VARCHAR(50);NOT NULL"`
	Alamat    string `gorm:"type:VARCHAR(100);NOT NULL"`
	Kecamatan string `gorm:"type:VARCHAR(50);NOT NULL"`
	Harga     string `gorm:"type:VARCHAR(20);NOT NULL"`
	Kapasitas string `gorm:"type:VARCHAR(20);NOT NULL"`
	Luas      string `gorm:"type:VARCHAR(15);NOT NULL"`
	Tag       string
	Fasilitas string
	Aturan    string `gorm:"many2many:gedung_aturan"`
	Booking   *[]Booking
}
type Tag struct {
	gorm.Model
	Tipe string `gorm:"type:VARCHAR(20);NOT NULL"`
}
