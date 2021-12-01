package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"together-backend/internal/database"
	"together-backend/internal/repositories"
	"together-backend/internal/usecases"

	"gorm.io/gorm"
)

var (
	db             *gorm.DB
	accountUseCase usecases.AccountUseCase
)

type ReqBodyLogin struct {
	Email    string
	Password string
}

type ReqBodyRegister struct {
	Name            string
	Email           string
	Password        string
	PasswordConfirm string `json:"password_confirm"`
}

func Test(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "test",
	})
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var reqBody ReqBodyLogin
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "request body is incorrected",
		})
		return
	}

	loginResponse, err := accountUseCase.Login(reqBody.Email, reqBody.Password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "email or password is incorrected",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "login successfully",
		"token":   loginResponse.Token,
		"user":    loginResponse.User,
	})
}

func Logout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "logout successfully",
	})
}

func Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var reqBody ReqBodyRegister
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": "request body is incorrected",
		})
		return
	}

	fmt.Println("reqBody: ", reqBody)

	registerResponse, err := accountUseCase.Register(reqBody.Name, reqBody.Email, reqBody.Password, reqBody.PasswordConfirm)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"message": err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "register successfully",
		"token":   registerResponse.Token,
		"user":    registerResponse.User,
	})
}

func init() {
	db = database.ConnectDB()
	accountRepo := repositories.NewUserRepo(db)
	accountUseCase = usecases.NewAccountUsecase(accountRepo)
}
