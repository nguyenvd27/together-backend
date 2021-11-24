package usecases

import (
	"fmt"
	"together-backend/internal/models"
	"together-backend/internal/repositories"
)

type CommentCase interface {
	GetCommentsByEventIdUsecase(eventId, page, size int) ([]models.Comment, int64, error)
	CreateCommentUsecase(reqBody *ReqBodyComment, eventId, userId int) (*models.Comment, error)
	DeleteCommentUsecase(commentId, eventId, userId int) (*models.Comment, error)
}

type commentUsecase struct {
	commentRepo   repositories.CommentRepo
	userEventRepo repositories.UserEventRepo
}

type ReqBodyComment struct {
	Content string
}

func NewCommentUsecase(commentRepo repositories.CommentRepo, userEventRepo repositories.UserEventRepo) CommentCase {
	return &commentUsecase{
		commentRepo:   commentRepo,
		userEventRepo: userEventRepo,
	}
}

func (uc *commentUsecase) CreateCommentUsecase(reqBody *ReqBodyComment, eventId, userId int) (*models.Comment, error) {
	if eventId <= 0 {
		return nil, fmt.Errorf("invalid event id")
	}
	if userId <= 0 {
		return nil, fmt.Errorf("invalid user id")
	}
	if reqBody.Content == "" {
		return nil, fmt.Errorf("content of comment is empty")
	}

	userEvent, err := uc.userEventRepo.GetUserFromEvent(userId, eventId)
	if err != nil && err.Error() != "record not found" {
		return nil, err
	}

	if userEvent == nil {
		return nil, fmt.Errorf("you haven't joined in the event yet")
	}

	newComment, err := uc.commentRepo.CreateComment(userId, eventId, reqBody.Content)
	if err != nil {
		return nil, err
	}

	return newComment, nil
}

func (uc *commentUsecase) GetCommentsByEventIdUsecase(eventId, page, size int) ([]models.Comment, int64, error) {

	if eventId <= 0 {
		return nil, int64(0), fmt.Errorf("invalid event id")
	}

	total, err := uc.commentRepo.CountCommentsByEventId(eventId)
	if err != nil {
		return nil, int64(0), err
	}

	comments, err := uc.commentRepo.GetCommentsByEventId(eventId, size, page)
	if err != nil {
		return nil, int64(0), err
	}

	return comments, total, nil
}

func (uc *commentUsecase) DeleteCommentUsecase(commentId, eventId, userId int) (*models.Comment, error) {
	if eventId <= 0 {
		return nil, fmt.Errorf("invalid event id")
	}
	if userId <= 0 {
		return nil, fmt.Errorf("invalid user id")
	}
	if commentId <= 0 {
		return nil, fmt.Errorf("invalid comment id")
	}

	userEvent, err := uc.userEventRepo.GetUserFromEvent(userId, eventId)
	if err != nil && err.Error() != "record not found" {
		return nil, err
	}

	if userEvent == nil {
		return nil, fmt.Errorf("you haven't joined in the event yet")
	}

	comment, err := uc.commentRepo.GetComment(commentId, userId, eventId)
	if err != nil {
		return nil, err
	}

	deleteComment, err := uc.commentRepo.DeleteComment(comment)
	if err != nil {
		return nil, err
	}

	return deleteComment, nil
}
