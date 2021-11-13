package handlers

import (
	"encoding/json"
	"mime/multipart"
	"net/http"
	"strconv"
	"time"
	"together-backend/internal/database"
	"together-backend/internal/repositories"
	"together-backend/internal/usecases"
)

var eventUsecase usecases.EventUseCase

func parseRequestBodyFromMultipartFrom(multipartFrom *multipart.Form) (*usecases.ReqBodyEvent, error) {
	var (
		reqBody usecases.ReqBodyEvent
		err     error
	)
	reqBody.Title = multipartFrom.Value["title"][0]
	reqBody.Content = multipartFrom.Value["content"][0]
	reqBody.CreatedBy, err = strconv.ParseUint(multipartFrom.Value["created_by"][0], 10, 64)
	if err != nil {
		return nil, err
	}
	reqBody.StartTime, err = time.Parse("2006-01-02T15:04:05Z0700", multipartFrom.Value["start_time"][0])
	if err != nil {
		return nil, err
	}
	reqBody.EndTime, err = time.Parse("2006-01-02T15:04:05Z0700", multipartFrom.Value["end_time"][0])
	if err != nil {
		return nil, err
	}
	reqBody.Location, err = strconv.Atoi(multipartFrom.Value["location"][0])
	if err != nil {
		return nil, err
	}
	reqBody.DetailLocation = multipartFrom.Value["detail_location"][0]
	return &reqBody, nil
}

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if err := r.ParseMultipartForm(32 << 20); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"message": err.Error(),
		})
		return
	}

	reqBody, err := parseRequestBodyFromMultipartFrom(r.MultipartForm)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"message": err.Error(),
		})
		return
	}

	files := r.MultipartForm.File["images"]
	imagesSlice, err := usecases.EventImageUpload(files)
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

func init() {
	db = database.ConnectDB()
	eventRepo := repositories.NewEventRepo(db)
	eventUsecase = usecases.NewEventUsecase(eventRepo)
}