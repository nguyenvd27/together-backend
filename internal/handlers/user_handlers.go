package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"together-backend/internal/database"
	"together-backend/internal/repositories"
	"together-backend/internal/usecases"

	"github.com/gorilla/mux"
)

var (
	userUsercase usecases.UserCase
)

// const SIZE int = 8

func GetUserDetail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	userId, err := strconv.Atoi(params["user_id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "failed to parse params",
		})
		return
	}

	user, err := userUsercase.GetUserDetailUsecase(userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"message": err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "get user detail successfully",
		"user":    user,
	})
}

func init() {
	db = database.ConnectDB()
	userRepo := repositories.NewUserRepo(db)
	userUsercase = usecases.NewUserUsecase(userRepo)
}
