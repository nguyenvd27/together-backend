package usecases

import (
	"fmt"
	"os"
	"time"
	"together-backend/internal/models"
	"together-backend/internal/repositories"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type AccountUseCase interface {
	Login(email, password string) (*AuthResponse, error)
	Register(name, email, password string) (*AuthResponse, error)
}

type accountUsecase struct {
	accountRepo repositories.UserRepo
}

type AuthResponse struct {
	Token string
	User  models.User
}

type Claims struct {
	UserId int    `json:"user_id"`
	Email  string `json:"email"`
	jwt.StandardClaims
}

func NewAccountUsecase(accountRepo repositories.UserRepo) AccountUseCase {
	return &accountUsecase{
		accountRepo: accountRepo,
	}
}

func (uc *accountUsecase) Login(email, password string) (*AuthResponse, error) {
	user, err := uc.accountRepo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(password)); err != nil {
		return nil, err
	}

	// Generate token
	claims := &Claims{
		UserId: int(user.Id),
		Email:  email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	var jwtKey = []byte(os.Getenv("SECRET_KEY"))
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return nil, err
	}

	return &AuthResponse{
		Token: tokenString,
		User:  user,
	}, nil
}

func (uc *accountUsecase) Register(name, email, password string) (*AuthResponse, error) {
	user, err := uc.accountRepo.GetUserByEmail(email)
	if err != nil && err.Error() != "record not found" {
		return nil, err
	}

	if user.Email == email {
		return nil, fmt.Errorf("email already exists")
	}

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)

	newUser, err := uc.accountRepo.CreateUser(name, email, hashPassword)
	if err != nil {
		return nil, err
	}

	// Generate token
	claims := &Claims{
		UserId: int(newUser.Id),
		Email:  email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	var jwtKey = []byte(os.Getenv("SECRET_KEY"))
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return nil, err
	}

	return &AuthResponse{
		Token: tokenString,
		User:  newUser,
	}, nil
}
