package repositories

import (
	"together-backend/internal/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserEventRepo interface {
	GetUserFromEvent(userId, eventId int) (*models.UserEvent, error)
	AddUserToEvent(userId, eventId int) (*models.UserEvent, error)
	RemoveUserFromEvent(userId, eventId int) (*models.UserEvent, error)
}

type userEventDB struct {
	db *gorm.DB
}

func NewUserEventRepo(db *gorm.DB) UserEventRepo {
	return &userEventDB{
		db: db,
	}
}

func (userEventDB *userEventDB) GetUserFromEvent(userId, eventId int) (*models.UserEvent, error) {
	var userEvent models.UserEvent
	err := userEventDB.db.Where("user_id = ? AND event_id = ?", userId, eventId).First(&userEvent).Error
	if err != nil {
		return nil, err
	}
	return &userEvent, nil
}

func (userEventDB *userEventDB) AddUserToEvent(userId, eventId int) (*models.UserEvent, error) {
	var userEvent models.UserEvent = models.UserEvent{
		UserId:  uint(userId),
		EventId: uint(eventId),
	}
	if err := userEventDB.db.Create(&userEvent).Error; err != nil {
		return nil, err
	}

	return &userEvent, nil
}

func (userEventDB *userEventDB) RemoveUserFromEvent(userId, eventId int) (*models.UserEvent, error) {

	var userEvent models.UserEvent = models.UserEvent{
		UserId:  uint(userId),
		EventId: uint(eventId),
	}
	if err := userEventDB.db.Clauses(clause.Returning{}).Delete(&userEvent).Error; err != nil {
		return nil, err
	}

	return &userEvent, nil
}
