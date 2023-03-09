package entity

import "gorm.io/gorm"

type Gedung struct {
	gorm.Model
	Nama      string      `gorm:"type:VARCHAR(50);NOT NULL"`
	Alamat    string      `gorm:"type:VARCHAR(100);NOT NULL"`
	Kecamatan string      `gorm:"type:VARCHAR(50);NOT NULL"`
	Harga     string      `gorm:"type:VARCHAR(20);NOT NULL"`
	Kapasitas string      `gorm:"type:VARCHAR(20);NOT NULL"`
	Luas      string      `gorm:"type:VARCHAR(15);NOT NULL"`
	Tag       []Tag       `gorm:"many2many:gedung_tag;"`
	Fasilitas []Fasilitas `gorm:"many2many:gedung_fasilitas"`
	Aturan    []Aturan    `gorm:"many2many:gedung_aturan"`
	Booking   []Booking
}
type Tag struct {
	gorm.Model
	Tipe string `gorm:"type:VARCHAR(20);NOT NULL"`
}
type Aturan struct {
	gorm.Model
	Aturan string `gorm:"type:VARCHAR(100);NOT NULL"`
}
type Fasilitas struct {
	gorm.Model
	Fasilitas string `gorm:"type:VARCHAR(100);NOT NULL"`
}
