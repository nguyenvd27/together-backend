package models

import (
	"time"

	"gorm.io/gorm"
)

type EventImage struct {
	Id        uint           `json:"id" gorm:"primaryKey"`
	EventId   uint           `json:"event_id"`
	ImageUrl  string         `json:"image_url"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
