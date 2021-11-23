package models

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	Id             uint           `json:"id" gorm:"primaryKey"`
	Title          string         `json:"title"`
	Content        string         `json:"content"`
	CreatedBy      uint64         `json:"created_by"`
	StartTime      time.Time      `json:"start_time"`
	EndTime        time.Time      `json:"end_time"`
	Location       int            `json:"location"`
	DetailLocation string         `json:"detail_location"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`
	EventImages    []EventImage   `json:"event_images" gorm:"foreignKey:EventId"`
	Comments       []Comment      `json:"comments" gorm:"foreignKey:EventId"`
	Users          []User         `json:"users" gorm:"many2many:user_events;"`
}
