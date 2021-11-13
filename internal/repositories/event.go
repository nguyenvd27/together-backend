package repositories

import (
	"fmt"
	"time"
	"together-backend/internal/models"

	"gorm.io/gorm"
)

type EventRepo interface {
	CreateEvent(title, content string, imageUrl []string, createdBy uint64, startTime, endTime time.Time, location int, detailLocation string) (*models.Event, error)
}

type eventDB struct {
	db *gorm.DB
}

func NewEventRepo(db *gorm.DB) EventRepo {
	return &eventDB{
		db: db,
	}
}

func imageUrls(imageUrl []string) []models.EventImage {
	var eventImageSlice []models.EventImage
	for _, value := range imageUrl {
		eventImageSlice = append(eventImageSlice, models.EventImage{
			ImageUrl: value,
		})
	}
	return eventImageSlice
}

func (eventDB *eventDB) CreateEvent(title, content string, imageUrl []string, createdBy uint64, startTime, endTime time.Time, location int, detailLocation string) (*models.Event, error) {
	event := models.Event{
		Title:          title,
		Content:        content,
		CreatedBy:      createdBy,
		StartTime:      startTime,
		EndTime:        endTime,
		Location:       location,
		DetailLocation: detailLocation,
		EventImages:    imageUrls(imageUrl),
	}
	if err := eventDB.db.Create(&event).Error; err != nil {
		return nil, fmt.Errorf("failed to create event")
	}
	return &event, nil
}
