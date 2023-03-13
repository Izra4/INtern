package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Nama     string `gorm:"NOT NULL"`
	Email    string `gorm:"unique;NOT NULL" json:"email"`
	Password string `gorm:"NOT NULL" json:"password"`
	Number   string `gorm:"type:varchar(20);NOT NULL" json:"number"`
	Bookings []Booking
	Payments []Payment
}
