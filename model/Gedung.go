package model

import (
	"InternBCC/database"
	"gorm.io/gorm"
)

type Gedung struct {
	gorm.Model
	Nama      string `json:"nama" binding:"required"`
	Alamat    string `json:"fasilitas" binding:"required"`
	Kecamatan string `json:"kecamatan" binding:"required"`
	Harga     string `json:"harga" binding:"required"`
	Kapasitas string `json:"kapasitas" binding:"required"`
	Luas      string `json:"luas" binding:"required"`
}

type name struct {
}

func GDummy() {
	G1 := Gedung{
		Model:     gorm.Model{},
		Nama:      "Graha Cakrawala",
		Alamat:    "Universitas Negeri Malang, Jl. Cakrawala",
		Kecamatan: "Lowokwaru",
		Harga:     "Rp. 50jt-200jt",
		Kapasitas: "100 Ribu Orang",
		Luas:      "2 km",
	}
	G2 := Gedung{
		Model:     gorm.Model{},
		Nama:      "Dome UMM",
		Alamat:    "Universitas Muhammadiyah Malang, Jl. Tlogomas",
		Kecamatan: "Dau",
		Harga:     "Rp. 50jt-200jt",
		Kapasitas: "50 Ribu Orang",
		Luas:      "1 km",
	}
	G3 := Gedung{
		Model:     gorm.Model{},
		Nama:      "Singhasari",
		Alamat:    "Jl. Singosari",
		Kecamatan: "Singosari",
		Harga:     "Rp. 25jt-100jt",
		Kapasitas: "75 Ribu Orang",
		Luas:      "2 km",
	}
	G4 := Gedung{
		Model:     gorm.Model{},
		Nama:      "Graha Tirta",
		Alamat:    "Jl. Bend. Sengguruh No.32",
		Kecamatan: "Lowokwaru",
		Harga:     "Rp. 25jt-50jt",
		Kapasitas: "25 Ribu Orang",
		Luas:      "1 km",
	}
	G5 := Gedung{
		Model:     gorm.Model{},
		Nama:      "Gedung Gajahmada",
		Alamat:    "Jl. Rampal Celaket",
		Kecamatan: "Klojen",
		Harga:     "Rp. 75jt-100jt",
		Kapasitas: "80 Ribu Orang",
		Luas:      "2 km",
	}
	G6 := Gedung{
		Model:     gorm.Model{},
		Nama:      "Graha Cakrawala",
		Alamat:    "Universitas Negeri Malang, Jl. Cakrawala",
		Kecamatan: "Lowokwaru",
		Harga:     "Rp. 50jt-200jt",
		Kapasitas: "100 Ribu Orang",
		Luas:      "2 km",
	}
	if err := database.DB.Create(&G1).Error; err != nil {
		return
	}
	if err := database.DB.Create(&G2).Error; err != nil {
		return
	}
	if err := database.DB.Create(&G3).Error; err != nil {
		return
	}
	if err := database.DB.Create(&G4).Error; err != nil {
		return
	}
	if err := database.DB.Create(&G5).Error; err != nil {
		return
	}
	if err := database.DB.Create(&G6).Error; err != nil {
		return
	}

	//if err := database.DB.Model(&G1).Association("Tag").Append([]Tag).Error(); err != nil {
	//	return
	//}
}
