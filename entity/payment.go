package entity

import (
	"time"
)

type Payment struct {
	ID        string    `gorm:"primarykey;type:VARCHAR(8);unique" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uint      `json:"user_id"`
	GedungID  uint      `json:"gedung_id"`
	Link      string    `gorm:"type:TEXT;NOT NULL" json:"link"`
	Nominal   int       `gorm:"type:int;NOT NULL" json:"nominal"`
}
