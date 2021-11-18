package usecases

import (
	"fmt"
	"time"
	"together-backend/internal/models"
	"together-backend/internal/repositories"
)

type EventUseCase interface {
	CreateEventUsecase(reqBody *ReqBodyEvent, imageUrl []string) (*models.Event, error)
	GetEventsUsecase(page, size int) ([]EventsCreatedByUser, int64, error)
	GetEventDetailUsecase(eventId int) (*EventsCreatedByUser, error)
	DeleteEventUsecase(eventId, userId int) (string, error)
}

type eventUsecase struct {
	eventRepo repositories.EventRepo
	userRepo  repositories.UserRepo
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

type EventsCreatedByUser struct {
	EventDetail   models.Event `json:"event_detail"`
	CreatedByUser models.User  `json:"created_by_user"`
}

func NewEventUsecase(eventRepo repositories.EventRepo, userRepo repositories.UserRepo) EventUseCase {
	return &eventUsecase{
		eventRepo: eventRepo,
		userRepo:  userRepo,
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

func (uc *eventUsecase) GetEventsUsecase(page, size int) ([]EventsCreatedByUser, int64, error) {
	var eventsCreatedByUsers []EventsCreatedByUser

	total, err := uc.eventRepo.CountEvents()
	if err != nil {
		return nil, int64(0), err
	}

	events, err := uc.eventRepo.GetEvents(page, size)
	if err != nil {
		return nil, int64(0), err
	}

	for i := 0; i < len(events); i++ {
		createdByUser, err := uc.userRepo.GetUserById(int64(events[i].CreatedBy))
		if err != nil {
			return nil, int64(0), fmt.Errorf("failed to get user who created the event")
		}
		eventsCreatedByUser := EventsCreatedByUser{
			EventDetail:   events[i],
			CreatedByUser: createdByUser,
		}
		eventsCreatedByUsers = append(eventsCreatedByUsers, eventsCreatedByUser)
	}

	return eventsCreatedByUsers, total, nil
}

func (uc *eventUsecase) GetEventDetailUsecase(eventId int) (*EventsCreatedByUser, error) {
	if eventId <= 0 {
		return nil, fmt.Errorf("invalid event id")
	}

	event, err := uc.eventRepo.GetEventDetail(eventId)
	if err != nil {
		return nil, err
	}

	createdByUser, err := uc.userRepo.GetUserById(int64(event.CreatedBy))
	if err != nil {
		return nil, err
	}

	eventsCreatedByUser := EventsCreatedByUser{
		EventDetail:   event,
		CreatedByUser: createdByUser,
	}
	return &eventsCreatedByUser, nil
}

func (uc *eventUsecase) DeleteEventUsecase(eventId, userId int) (string, error) {
	if eventId <= 0 {
		return "", fmt.Errorf("invalid event id")
	}
	if userId <= 0 {
		return "", fmt.Errorf("invalid user id")
	}

	event, err := uc.eventRepo.GetEventByEventIdAndCreatedBy(eventId, userId)
	if err != nil {
		return "", err
	}

	fmt.Println("event will delete: ", event)

	mess, err := uc.eventRepo.DeleteEvent(event)
	if err != nil {
		return "", err
	}

	return mess, nil
}
