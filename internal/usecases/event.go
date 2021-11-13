package usecases

import (
	"fmt"
	"time"
	"together-backend/internal/models"
	"together-backend/internal/repositories"
)

type EventUseCase interface {
	CreateEventUsecase(reqBody *ReqBodyEvent, imageUrl []string) (*models.Event, error)
}

type eventUsecase struct {
	eventRepo repositories.EventRepo
}

type ReqBodyEvent struct {
	Title          string
	Content        string
	CreatedBy      uint64
	StartTime      time.Time
	EndTime        time.Time
	Location       int
	DetailLocation string
}

func NewEventUsecase(eventRepo repositories.EventRepo) EventUseCase {
	return &eventUsecase{
		eventRepo: eventRepo,
	}
}

func (uc *eventUsecase) CreateEventUsecase(reqBody *ReqBodyEvent, imageUrl []string) (*models.Event, error) {
	if reqBody.Title == "" {
		return nil, fmt.Errorf("title cannot be empty")
	}
	if reqBody.Content == "" {
		return nil, fmt.Errorf("content cannot be empty")
	}
	if reqBody.CreatedBy == uint64(0) {
		return nil, fmt.Errorf("created_by cannot be empty")
	}
	newEvent, err := uc.eventRepo.CreateEvent(reqBody.Title, reqBody.Content, imageUrl, reqBody.CreatedBy, reqBody.StartTime, reqBody.EndTime, reqBody.Location, reqBody.DetailLocation)
	if err != nil {
		return nil, err
	}
	return newEvent, nil
}
