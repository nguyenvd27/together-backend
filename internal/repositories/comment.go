package repositories

import (
	"together-backend/internal/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CommentRepo interface {
	CreateComment(userId, eventId int, content string) (*models.Comment, error)
	GetCommentsByEventId(eventId, size, page int) ([]models.Comment, error)
	CountCommentsByEventId(eventId int) (int64, error)
	DeleteComment(comment *models.Comment) (*models.Comment, error)
	GetComment(commentId, userId, eventId int) (*models.Comment, error)
}

type commentDB struct {
	db *gorm.DB
}

func NewCommentRepo(db *gorm.DB) CommentRepo {
	return &commentDB{
		db: db,
	}
}

func (commentDB *commentDB) CreateComment(userId, eventId int, content string) (*models.Comment, error) {
	comment := models.Comment{
		EventId: uint(eventId),
		UserId:  uint(userId),
		Content: content,
	}
	if err := commentDB.db.Create(&comment).Error; err != nil {
		return nil, err
	}
	commentDB.db.Preload("User").Find(&comment)

	return &comment, nil
}

func (commentDB *commentDB) GetCommentsByEventId(eventId, size, page int) ([]models.Comment, error) {
	var comments []models.Comment
	err := commentDB.db.Preload("User").
		Where("event_id = ?", eventId).
		Limit(size).Offset((page - 1) * size).
		Order("created_at desc").
		Find(&comments).Error
	if err != nil {
		return nil, err
	}

	return comments, nil
}

func (commentDB *commentDB) CountCommentsByEventId(eventId int) (int64, error) {
	var (
		total    int64
		comments models.Comment
	)
	err := commentDB.db.Where("event_id = ?", eventId).Find(&comments).Count(&total).Error
	if err != nil {
		return int64(0), err
	}

	return total, nil
}

func (commentDB *commentDB) GetComment(commentId, userId, eventId int) (*models.Comment, error) {
	var comment models.Comment
	err := commentDB.db.Preload("User").
		Where("id = ? AND user_id = ? AND event_id = ?", commentId, userId, eventId).
		First(&comment).Error
	if err != nil {
		return nil, err
	}

	return &comment, nil
}

func (commentDB *commentDB) DeleteComment(comment *models.Comment) (*models.Comment, error) {
	if err := commentDB.db.Clauses(clause.Returning{}).
		Unscoped().Delete(&comment).Error; err != nil {
		return nil, err
	}

	return comment, nil
}
