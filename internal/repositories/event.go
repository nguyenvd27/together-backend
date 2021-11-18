package repositories

import (
	"fmt"
	"time"
	"together-backend/internal/models"

	"gorm.io/gorm"
)

type EventRepo interface {
	CreateEvent(title, content string, imageUrl []string, createdBy uint64, startTime, endTime time.Time, location int, detailLocation string) (*models.Event, error)
	GetEvents(page, size int) ([]models.Event, error)
	CountEvents() (int64, error)
	GetEventDetail(eventId int) (models.Event, error)
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

func (eventDB *eventDB) GetEvents(page, size int) ([]models.Event, error) {
	var events []models.Event
	err := eventDB.db.Preload("Users").Preload("EventImages").
		Limit(size).Offset((page - 1) * size).
		Order("created_at desc").
		Find(&events).Error
	if err != nil {
		return nil, fmt.Errorf("failed to find events")
	}

	return events, nil
}

func (eventDB *eventDB) CountEvents() (int64, error) {
	var total int64

	err := eventDB.db.Table("events").Count(&total).Error
	if err != nil {
		return int64(0), fmt.Errorf("failed to count events")
	}

	return total, nil
}

func (eventDB *eventDB) GetEventDetail(eventId int) (models.Event, error) {
	var event models.Event
	err := eventDB.db.Preload("Users").Preload("EventImages").
		First(&event, eventId).Error
	if err != nil {
		return models.Event{}, err
	}

	return event, nil
}
