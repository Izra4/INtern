package entity

import "gorm.io/gorm"

type Payment struct {
	gorm.Model
	UserID   uint
	GedungID uint
	Link     string
}
