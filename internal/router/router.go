package router

import (
	"together-backend/internal/handlers"

	"together-backend/internal/middleware"

	"github.com/gorilla/mux"
)

func New() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", handlers.Test).Methods("GET")

	router.HandleFunc("/api/v1/login", handlers.Login).Methods("POST")
	router.HandleFunc("/api/v1/register", handlers.Register).Methods("POST")
	router.HandleFunc("/api/v1/logout", middleware.Auth(handlers.Logout)).Methods("POST")

	router.HandleFunc("/api/v1/events", middleware.Auth(handlers.CreateEvent)).Methods("POST")
	router.HandleFunc("/api/v1/events", handlers.GetEvents).Methods("GET")
	router.HandleFunc("/api/v1/events/{event_id}", handlers.GetEventDetail).Methods("GET")
	router.HandleFunc("/api/v1/events/{event_id}", middleware.Auth(handlers.DeleteEvent)).Methods("DELETE")
	router.HandleFunc("/api/v1/events/{event_id}", middleware.Auth(handlers.UpdateEvent)).Methods("PUT")
	router.HandleFunc("/api/v1/events/{event_id}/join", middleware.Auth(handlers.JoinEvent)).Methods("POST")

	router.HandleFunc("/api/v1/events/{event_id}/comments", handlers.GetCommentsByEventId).Methods("GET")
	router.HandleFunc("/api/v1/events/{event_id}/comments", middleware.Auth(handlers.CreateComment)).Methods("POST")
	router.HandleFunc("/api/v1/events/{event_id}/comments/{comment_id}", middleware.Auth(handlers.DeleteComment)).Methods("DELETE")

	router.HandleFunc("/api/v1/users/{user_id}", middleware.Auth(handlers.GetUserDetail)).Methods("GET")
	router.HandleFunc("/api/v1/users/{user_id}", middleware.Auth(handlers.UpdateProfile)).Methods("PUT")
	router.HandleFunc("/api/v1/users/{user_id}/change_password", middleware.Auth(handlers.ChangePassword)).Methods("PUT")

	return router
}
