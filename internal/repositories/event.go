package repositories

import (
	"time"
	"together-backend/internal/models"

	"gorm.io/gorm"
)

type EventRepo interface {
	CreateEvent(title, content string, imageUrl []string, createdBy uint64, startTime, endTime time.Time, location int, detailLocation string) (*models.Event, error)
	GetEvents(page, size int) ([]models.Event, error)
	CountEvents() (int64, error)
	GetEventDetail(eventId int) (models.Event, error)
	GetEventByEventIdAndCreatedBy(eventId, createdBy int) (models.Event, error)
	DeleteEvent(event models.Event) (string, error)
	UpdateEvent(event models.Event, title, content string, imageUrl []string, startTime, endTime time.Time, location int, detailLocation string) (*models.Event, error)
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
		return nil, err
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
		return nil, err
	}

	return events, nil
}

func (eventDB *eventDB) CountEvents() (int64, error) {
	var (
		total  int64
		events models.Event
	)
	err := eventDB.db.Find(&events).Count(&total).Error
	if err != nil {
		return int64(0), err
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

func (eventDB *eventDB) GetEventByEventIdAndCreatedBy(eventId, createdBy int) (models.Event, error) {
	var event models.Event
	err := eventDB.db.Preload("Users").Preload("EventImages").
		First(&event, "id = ? AND created_by = ?", eventId, createdBy).Error
	if err != nil {
		return models.Event{}, err
	}

	return event, nil
}

func (eventDB *eventDB) DeleteEvent(event models.Event) (string, error) {
	err := eventDB.db.Delete(&event).Error
	if err != nil {
		return "", err
	}

	return "deleted successfully", nil
}

func (eventDB *eventDB) UpdateEvent(event models.Event, title, content string, imageUrl []string, startTime, endTime time.Time, location int, detailLocation string) (*models.Event, error) {
	editEvent := models.Event{
		Id:             event.Id,
		Title:          title,
		Content:        content,
		StartTime:      startTime,
		EndTime:        endTime,
		Location:       location,
		DetailLocation: detailLocation,
		EventImages:    imageUrls(imageUrl),
	}

	if len(event.EventImages) > 0 {
		err := eventDB.db.Model(&event.EventImages).Delete(event.EventImages).Error
		if err != nil {
			return nil, err
		}
	}

	if err := eventDB.db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&editEvent).Error; err != nil {
		return nil, err
	}

	return &event, nil
}
