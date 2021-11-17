package models

type UserEvent struct {
	UserId  uint `gorm:"primaryKey" column:"user_id"`
	EventId uint `gorm:"primaryKey" column:"event_id"`
}
