package model

import (
	"InternBCC/database"
	"gorm.io/gorm"
)

type Aturan struct {
	gorm.Model
	Aturan string `json:"aturan" binding:"required"`
}

func ADummy() {
	a1 := Aturan{Aturan: "jangan bawa miras"}
	a2 := Aturan{Aturan: "Tidak boleh membawa property"}
	a3 := Aturan{Aturan: "Dilarang merusak property"}
	a4 := Aturan{Aturan: "Tata ulang tempat setelah selesai"}

	if err := database.DB.Create(&a1).Error; err != nil {
		return
	}
	if err := database.DB.Create(&a2).Error; err != nil {
		return
	}
	if err := database.DB.Create(&a3).Error; err != nil {
		return
	}
	if err := database.DB.Create(&a4).Error; err != nil {
		return
	}

}
