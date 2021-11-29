package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"together-backend/internal/database"
	"together-backend/internal/repositories"
	"together-backend/internal/transfers"
	"together-backend/internal/usecases"

	"github.com/gorilla/mux"
)

var (
	userUsercase usecases.UserCase
)

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

func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if err := r.ParseMultipartForm(32 << 20); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"message": err.Error(),
		})
		return
	}

	userId := r.Context().Value("currentUserID").(int)

	params := mux.Vars(r)
	userIdParam, err := strconv.Atoi(params["user_id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "failed to parse params",
		})
		return
	}

	if userId != userIdParam {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "you don't have permission to update profile for this account",
		})
		return
	}

	reqBody, err := transfers.ParseRequestUpdateProfile(r.MultipartForm)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"message": err.Error(),
		})
		return
	}

	var avatarUrl string
	if len(r.MultipartForm.File["avatar"]) > 0 {
		file := r.MultipartForm.File["avatar"][0]
		avatarUrl, err = uploadUsecase.UploadAvatar(file)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{
				"message": err.Error(),
			})
			return
		}
	}

	updatedUser, err := userUsercase.UpdateProfilelUsecase(userId, reqBody.Name, reqBody.Address, avatarUrl)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"message": err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "update profile successfully",
		"user":    updatedUser,
	})
}

type reqBodyChangePassword struct {
	OldPassword     string `json:"old_password"`
	NewPassword     string `json:"new_password"`
	PasswordConfirm string `json:"password_confirm"`
}

func ChangePassword(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userId := r.Context().Value("currentUserID").(int)

	params := mux.Vars(r)
	userIdParam, err := strconv.Atoi(params["user_id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "failed to parse params",
		})
		return
	}

	if userId != userIdParam {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "you don't have permission to change password for this account",
		})
		return
	}

	var reqBody reqBodyChangePassword
	err = json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"message": err.Error(),
		})
		return
	}

	updatedUser, err := userUsercase.ChangePasswordUsecase(userId, reqBody.OldPassword, reqBody.NewPassword, reqBody.PasswordConfirm)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"message": err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "change password successfully",
		"user":    updatedUser,
	})
}

func init() {
	db = database.ConnectDB()
	userRepo := repositories.NewUserRepo(db)
	userUsercase = usecases.NewUserUsecase(userRepo)
}
