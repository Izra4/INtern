package entity

import "gorm.io/gorm"

type Link struct {
	gorm.Model
	Link     string `gorm:"type:TEXT;NOT NULL"`
	GedungID uint
}
