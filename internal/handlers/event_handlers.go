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
	eventUsecase  usecases.EventUseCase
	uploadUsecase usecases.UploadUseCase
)

const SIZE_PER_PAGE int = 8

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if err := r.ParseMultipartForm(32 << 20); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"message": err.Error(),
		})
		return
	}

	reqBody, err := transfers.ParseRequestBodyFromMultipartFrom(r.MultipartForm)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"message": err.Error(),
		})
		return
	}

	files := r.MultipartForm.File["images"]
	imagesSlice, err := uploadUsecase.EventImageUpload(files)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"message": err.Error(),
		})
		return
	}

	newEvent, err := eventUsecase.CreateEventUsecase(reqBody, imagesSlice)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"message": err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "create event successfully",
		"event":   newEvent,
	})
}

func GetEvents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var (
		err  error
		page int = 1
	)
	queries := r.URL.Query()
	if len(queries["page"]) > 0 {
		page, err = strconv.Atoi(queries["page"][0])
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "failed to parse query page",
			})
			return
		}
	}

	events, total, err := eventUsecase.GetEventsUsecase(page, SIZE_PER_PAGE)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"message": err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "get events successfully",
		"events":  events,
		"total":   total,
		"page":    page,
	})
}

func GetEventDetail(w http.ResponseWriter, r *http.Request) {
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

	event, err := eventUsecase.GetEventDetailUsecase(eventId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"message": err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "get events successfully",
		"event":   event,
	})
}

func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	eventId, err := strconv.Atoi(params["event_id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "failed to parse event id",
		})
		return
	}

	userId := r.Context().Value("currentUserID").(int)

	mess, err := eventUsecase.DeleteEventUsecase(eventId, userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"message": err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": mess,
	})
}

func UpdateEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if err := r.ParseMultipartForm(32 << 20); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"message": err.Error(),
		})
		return
	}

	userId := r.Context().Value("currentUserID").(int)

	reqBody, err := transfers.ParseRequestUpdateEvent(r.MultipartForm)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"message": err.Error(),
		})
		return
	}

	files := r.MultipartForm.File["images"]
	imagesSlice, err := uploadUsecase.EventImageUpload(files)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"message": err.Error(),
		})
		return
	}

	updatedEvent, err := eventUsecase.UpdateEventUsecase(userId, reqBody, imagesSlice)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"message": err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "update event successfully",
		"event":   updatedEvent,
	})
}

func JoinEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	userId := r.Context().Value("currentUserID").(int)

	params := mux.Vars(r)
	eventId, err := strconv.Atoi(params["event_id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "failed to parse params",
		})
		return
	}

	event, mess, err := eventUsecase.JoinEventUsecase(userId, eventId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"message": err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": mess,
		"event":   event,
	})
}

func init() {
	db = database.ConnectDB()
	eventRepo := repositories.NewEventRepo(db)
	userRepo := repositories.NewUserRepo(db)
	imageRepo := repositories.NewImageRepo(db)
	userEventRepo := repositories.NewUserEventRepo(db)
	eventUsecase = usecases.NewEventUsecase(eventRepo, userRepo, imageRepo, userEventRepo)
	uploadUsecase = usecases.NewUploadUsecase(imageRepo)
}
