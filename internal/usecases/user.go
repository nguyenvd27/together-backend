package usecases

import (
	"fmt"
	"together-backend/internal/models"
	"together-backend/internal/repositories"
)

type UserCase interface {
	GetUserDetailUsecase(userId int) (*models.User, error)
}

type userUsecase struct {
	userRepo repositories.UserRepo
}

// type ReqBodyComment struct {
// 	Content string
// }

func NewUserUsecase(userRepo repositories.UserRepo) UserCase {
	return &userUsecase{
		userRepo: userRepo,
	}
}

func (uc *userUsecase) GetUserDetailUsecase(userId int) (*models.User, error) {

	if userId <= 0 {
		return nil, fmt.Errorf("invalid user id")
	}

	user, err := uc.userRepo.GetUserDetail(userId)
	if err != nil {
		return nil, err
	}

	return user, nil
}
