package usecases

import (
	"fmt"
	"time"
	"together-backend/internal/models"
	"together-backend/internal/repositories"
)

type EventUseCase interface {
	CreateEventUsecase(reqBody *ReqBodyEvent, imageUrl []string) (*models.Event, error)
	GetEventsUsecase(page, size, userId int, search, qType string) ([]EventsCreatedByUser, int64, error)
	GetEventDetailUsecase(eventId int) (*EventsCreatedByUser, error)
	DeleteEventUsecase(eventId, userId int) (string, error)
	UpdateEventUsecase(userId int, reqBody *ReqBodyEditEvent, imageUrl []string) (*models.Event, error)
	JoinEventUsecase(userId, eventId int) (*EventsCreatedByUser, string, error)
}

type eventUsecase struct {
	eventRepo     repositories.EventRepo
	userRepo      repositories.UserRepo
	imageRepo     repositories.ImageRepo
	userEventRepo repositories.UserEventRepo
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

type ReqBodyEditEvent struct {
	Id             uint64
	Title          string
	Content        string
	CreatedBy      uint64
	StartTime      time.Time
	EndTime        time.Time
	Location       int
	DetailLocation string
}

func NewEventUsecase(eventRepo repositories.EventRepo, userRepo repositories.UserRepo, imageRepo repositories.ImageRepo, userEventRepo repositories.UserEventRepo) EventUseCase {
	return &eventUsecase{
		eventRepo:     eventRepo,
		userRepo:      userRepo,
		imageRepo:     imageRepo,
		userEventRepo: userEventRepo,
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
	if reqBody.StartTime.After(reqBody.EndTime) {
		return nil, fmt.Errorf("start time must be less than end time")
	}
	currentTime := time.Now()
	if currentTime.After(reqBody.EndTime) {
		return nil, fmt.Errorf("end time must be greater than current time")
	}
	newEvent, err := uc.eventRepo.CreateEvent(reqBody.Title, reqBody.Content, imageUrl, reqBody.CreatedBy, reqBody.StartTime, reqBody.EndTime, reqBody.Location, reqBody.DetailLocation)
	if err != nil {
		return nil, err
	}
	return newEvent, nil
}

func (uc *eventUsecase) GetEventsUsecase(page, size, userId int, search, qType string) ([]EventsCreatedByUser, int64, error) {
	var eventsCreatedByUsers []EventsCreatedByUser

	total, err := uc.eventRepo.CountEvents(userId, search, qType)
	if err != nil {
		return nil, int64(0), err
	}

	events, err := uc.eventRepo.GetEvents(page, size, userId, search, qType)
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

func (uc *eventUsecase) UpdateEventUsecase(userId int, reqBody *ReqBodyEditEvent, imageUrl []string) (*models.Event, error) {
	if reqBody.Title == "" {
		return nil, fmt.Errorf("title cannot be empty")
	}
	if reqBody.Content == "" {
		return nil, fmt.Errorf("content cannot be empty")
	}
	if reqBody.CreatedBy == uint64(0) {
		return nil, fmt.Errorf("created_by cannot be empty")
	}
	if reqBody.StartTime.After(reqBody.EndTime) {
		return nil, fmt.Errorf("start time must be less than end time")
	}
	event, err := uc.eventRepo.GetEventByEventIdAndCreatedBy(int(reqBody.Id), userId)
	if err != nil {
		return nil, err
	}

	updatedEvent, err := uc.eventRepo.UpdateEvent(event, reqBody.Title, reqBody.Content, imageUrl, reqBody.StartTime, reqBody.EndTime, reqBody.Location, reqBody.DetailLocation)
	if err != nil {
		return nil, err
	}
	return updatedEvent, nil
}

func (uc *eventUsecase) JoinEventUsecase(userId, eventId int) (*EventsCreatedByUser, string, error) {

	var (
		event models.Event
		mess  string
	)

	userEventGet, err := uc.userEventRepo.GetUserFromEvent(userId, eventId)
	if err != nil && err.Error() != "record not found" {
		return nil, "", err
	}

	if userEventGet != nil {
		userEventRemove, err := uc.userEventRepo.RemoveUserFromEvent(userId, eventId)
		if err != nil {
			return nil, "", err
		}
		event, err = uc.eventRepo.GetEventDetail(int(userEventRemove.EventId))
		if err != nil {
			return nil, "", err
		}
		mess = "removed from the event successfully"
	} else {
		userEventAdd, err := uc.userEventRepo.AddUserToEvent(userId, eventId)
		if err != nil {
			return nil, "", err
		}
		event, err = uc.eventRepo.GetEventDetail(int(userEventAdd.EventId))
		if err != nil {
			return nil, "", err
		}
		mess = "joined the event successfully"
	}

	createdByUser, err := uc.userRepo.GetUserById(int64(event.CreatedBy))
	if err != nil {
		return nil, "", err
	}

	eventsCreatedByUser := EventsCreatedByUser{
		EventDetail:   event,
		CreatedByUser: createdByUser,
	}
	return &eventsCreatedByUser, mess, nil
}
