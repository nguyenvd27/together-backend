package usecases

import (
	"fmt"
	"os"
	"time"
	"together-backend/internal/models"
	"together-backend/internal/repositories"
	"together-backend/pkg"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type AccountUseCase interface {
	Login(email, password string) (*AuthResponse, error)
	Register(name, email, password, passwordConfirm string) (*AuthResponse, error)
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
	if !pkg.ValidateEmail(email) {
		return nil, fmt.Errorf("email is not valid")
	}
	if len(password) == 0 {
		return nil, fmt.Errorf("password is empty")
	}

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
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
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

func (uc *accountUsecase) Register(name, email, password, passwordConfirm string) (*AuthResponse, error) {
	if name == "" {
		return nil, fmt.Errorf("name is empty")
	}
	if !pkg.ValidateEmail(email) {
		return nil, fmt.Errorf("email is not valid")
	}
	if len(password) < 8 {
		return nil, fmt.Errorf("password must be at least 8 characters")
	}
	if password != passwordConfirm {
		return nil, fmt.Errorf("password and confirm password does not match")
	}

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
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
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
