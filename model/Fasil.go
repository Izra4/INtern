package model

import (
	"InternBCC/database"
	"gorm.io/gorm"
)

type Fasilitas struct {
	gorm.Model
	Fasilitas string `json:"fasilitas" binding:"required"`
}

func FasDummy() {
	f1 := Fasilitas{
		Fasilitas: "bangku",
	}
	f2 := Fasilitas{
		Fasilitas: "meja",
	}
	f3 := Fasilitas{
		Fasilitas: "taplak",
	}
	f4 := Fasilitas{
		Fasilitas: "kaca",
	}
	if err := database.DB.Create(&f1).Error; err != nil {
		return
	}
	if err := database.DB.Create(&f2).Error; err != nil {
		return
	}
	if err := database.DB.Create(&f3).Error; err != nil {
		return
	}
	if err := database.DB.Create(&f4).Error; err != nil {
		return
	}

}
