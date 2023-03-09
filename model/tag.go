package model

import (
	"InternBCC/database"
	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	Tipe string `json:"tipe" binding:"required"`
}

func TagDummy() {
	tag1 := Tag{Tipe: "Pernikahan"}
	tag2 := Tag{Tipe: "Konser"}
	tag3 := Tag{Tipe: "Seminar"}
	tag4 := Tag{Tipe: "Gathering"}
	tag5 := Tag{Tipe: "Wisuda"}
	if err := database.DB.Create(&tag1).Error; err != nil {
		return
	}
	if err := database.DB.Create(&tag2).Error; err != nil {
		return
	}
	if err := database.DB.Create(&tag3).Error; err != nil {
		return
	}
	if err := database.DB.Create(&tag4).Error; err != nil {
		return
	}
	if err := database.DB.Create(&tag5).Error; err != nil {
		return
	}
}
