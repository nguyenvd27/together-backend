package repositories

import (
	"together-backend/internal/models"

	"gorm.io/gorm"
)

type ImageRepo interface {
	DeleteImageByEventId(eventId int) (string, error)
	GetImageByUrl(url string) (models.EventImage, error)
}

type imageDB struct {
	db *gorm.DB
}

func NewImageRepo(db *gorm.DB) ImageRepo {
	return &imageDB{
		db: db,
	}
}

func (imageDB *imageDB) GetImageByUrl(url string) (models.EventImage, error) {
	var image models.EventImage
	err := imageDB.db.Where("image_url = ?", url).First(&image).Error

	return image, err
}

func (imageDB *imageDB) DeleteImageByEventId(eventId int) (string, error) {
	var image models.EventImage
	err := imageDB.db.Where("event_id = ?", eventId).Unscoped().Delete(&image).Error
	if err != nil {
		return "", err
	}
	return "deleted successfully", nil
}
