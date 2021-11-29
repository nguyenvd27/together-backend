package repositories

import (
	"together-backend/internal/models"

	"gorm.io/gorm"
)

type UserRepo interface {
	GetUserByEmail(email string) (models.User, error)
	CreateUser(name, email string, password []byte) (models.User, error)
	GetUserById(id int64) (models.User, error)
	GetUserDetail(id int) (*models.User, error)
	UpdateProfile(user *models.User, name string, address int) (*models.User, error)
	UpdateProfileWithAvatar(user *models.User, name string, address int, avatarUrl string) (*models.User, error)
	ChangePassword(user *models.User, hashPassword []byte) (*models.User, error)
}

type userDB struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &userDB{
		db: db,
	}
}

func (userDB *userDB) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	err := userDB.db.Where("email = ?", email).First(&user).Error

	return user, err
}

func (userDB *userDB) CreateUser(name, email string, password []byte) (models.User, error) {
	var user = models.User{
		Name:     name,
		Email:    email,
		Password: password,
	}

	err := userDB.db.Create(&user).Error

	return user, err
}

func (userDB *userDB) GetUserById(id int64) (models.User, error) {
	var user models.User
	err := userDB.db.First(&user, id).Error

	return user, err
}

func (userDB *userDB) GetUserDetail(id int) (*models.User, error) {
	var user models.User
	err := userDB.db.Preload("Events.Users").Preload("Events.EventImages").Preload("Events").First(&user, id).Error
	if err != nil {
		return nil, err
	}

	return &user, err
}

func (userDB *userDB) UpdateProfile(user *models.User, name string, address int) (*models.User, error) {
	err := userDB.db.Model(user).Updates(map[string]interface{}{"name": name, "address": address}).Error
	if err != nil {
		return nil, err
	}

	return user, err
}

func (userDB *userDB) UpdateProfileWithAvatar(user *models.User, name string, address int, avatarUrl string) (*models.User, error) {
	err := userDB.db.Model(user).Updates(map[string]interface{}{"name": name, "address": address, "avatar": avatarUrl}).Error
	if err != nil {
		return nil, err
	}

	return user, err
}

func (userDB *userDB) ChangePassword(user *models.User, hashPassword []byte) (*models.User, error) {
	err := userDB.db.Model(user).Updates(map[string]interface{}{"password": hashPassword}).Error
	if err != nil {
		return nil, err
	}

	return user, err
}
