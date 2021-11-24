package models

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	Id      uint `json:"id" gorm:"primaryKey"`
	EventId uint `json:"event_id"`
	// Event     Event          `gorm:"references:Id"`
	UserId    uint           `json:"user_id"`
	User      User           `json:"user" gorm:"references:Id"`
	Content   string         `json:"content"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
