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
