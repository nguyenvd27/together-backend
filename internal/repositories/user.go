package repositories

import (
	"together-backend/internal/models"

	"gorm.io/gorm"
)

type UserRepo interface {
	GetUserByEmail(email string) (models.User, error)
	CreateUser(name, email string, password []byte) (models.User, error)
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
