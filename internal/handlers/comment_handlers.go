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
	commentUsercase usecases.CommentCase
)

const SIZE int = 8

func GetCommentsByEventId(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var (
		err         error
		commentPage int = 1
	)
	queries := r.URL.Query()
	if len(queries["comment_page"]) > 0 {
		commentPage, err = strconv.Atoi(queries["comment_page"][0])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "failed to parse comment page",
			})
			return
		}
	}

	params := mux.Vars(r)
	eventId, err := strconv.Atoi(params["event_id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "failed to parse params",
		})
		return
	}

	comments, total, err := commentUsercase.GetCommentsByEventIdUsecase(eventId, commentPage, SIZE)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"message": err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":      "get comments successfully",
		"comments":     comments,
		"event_id":     eventId,
		"total":        total,
		"comment_page": commentPage,
	})
}

func CreateComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var reqBodyComment *usecases.ReqBodyComment
	err := json.NewDecoder(r.Body).Decode(&reqBodyComment)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "failed to parse request body",
		})
		return
	}

	params := mux.Vars(r)
	eventId, err := strconv.Atoi(params["event_id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "failed to parse params",
		})
		return
	}

	userId := r.Context().Value("currentUserID").(int)

	newComment, err := commentUsercase.CreateCommentUsecase(reqBodyComment, eventId, userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"message": err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":  "created new comment successfully",
		"comment":  newComment,
		"event_id": eventId,
	})
}

func DeleteComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	eventId, err := strconv.Atoi(params["event_id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "failed to parse params",
		})
		return
	}

	commentId, err := strconv.Atoi(params["comment_id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "failed to parse params",
		})
		return
	}

	userId := r.Context().Value("currentUserID").(int)

	deleteComment, err := commentUsercase.DeleteCommentUsecase(commentId, eventId, userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"message": err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message":  "deleted comment successfully",
		"comment":  deleteComment,
		"event_id": eventId,
	})
}

func init() {
	db = database.ConnectDB()
	commentRepo := repositories.NewCommentRepo(db)
	userEventRepo := repositories.NewUserEventRepo(db)

	commentUsercase = usecases.NewCommentUsecase(commentRepo, userEventRepo)
}
