package usecases

import (
	"fmt"
	"together-backend/internal/models"
	"together-backend/internal/repositories"

	"golang.org/x/crypto/bcrypt"
)

type UserCase interface {
	GetUserDetailUsecase(userId int) (*models.User, error)
	UpdateProfilelUsecase(userId int, name string, address int, avatarUrl string) (*models.User, error)
	ChangePasswordUsecase(userId int, oldPassword, newPassword, passwordConfirm string) (*models.User, error)
}

type userUsecase struct {
	userRepo repositories.UserRepo
}

type ReqBodyUpdateProfile struct {
	Name    string
	Address int
}

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

func (uc *userUsecase) UpdateProfilelUsecase(userId int, name string, address int, avatarUrl string) (*models.User, error) {

	var (
		user        *models.User
		updatedUser *models.User
		err         error
	)
	if userId <= 0 {
		return nil, fmt.Errorf("invalid user id")
	}
	if name == "" {
		return nil, fmt.Errorf("name cannot be empty")
	}
	if address <= 0 {
		return nil, fmt.Errorf("invalid address")
	}

	user, err = uc.userRepo.GetUserDetail(userId)
	if err != nil {
		return nil, err
	}

	if avatarUrl == "" {
		updatedUser, err = uc.userRepo.UpdateProfile(user, name, address)
		if err != nil {
			return nil, err
		}
	} else {
		updatedUser, err = uc.userRepo.UpdateProfileWithAvatar(user, name, address, avatarUrl)
		if err != nil {
			return nil, err
		}
	}

	return updatedUser, nil
}

func (uc *userUsecase) ChangePasswordUsecase(userId int, oldPassword, newPassword, passwordConfirm string) (*models.User, error) {

	var (
		user        *models.User
		updatedUser *models.User
		err         error
	)
	if userId <= 0 {
		return nil, fmt.Errorf("invalid user id")
	}
	if len(newPassword) < 8 {
		return nil, fmt.Errorf("password must be at least 8 characters")
	}
	if newPassword != passwordConfirm {
		return nil, fmt.Errorf("password and confirm password does not match")
	}

	user, err = uc.userRepo.GetUserDetail(userId)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(oldPassword)); err != nil {
		return nil, fmt.Errorf("incorrect old password")
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), 14)
	if err != nil {
		return nil, err
	}

	updatedUser, err = uc.userRepo.ChangePassword(user, hashPassword)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}
