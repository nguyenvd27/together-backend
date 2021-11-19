package router

import (
	"together-backend/internal/handlers"

	"together-backend/internal/middleware"

	"github.com/gorilla/mux"
)

func New() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/api/v1/test", middleware.Auth(handlers.Test)).Methods("GET")

	router.HandleFunc("/api/v1/login", handlers.Login).Methods("POST")
	router.HandleFunc("/api/v1/register", handlers.Register).Methods("POST")
	router.HandleFunc("/api/v1/logout", middleware.Auth(handlers.Logout)).Methods("POST")

	router.HandleFunc("/api/v1/events", middleware.Auth(handlers.CreateEvent)).Methods("POST")
	router.HandleFunc("/api/v1/events", middleware.Auth(handlers.GetEvents)).Methods("GET")
	router.HandleFunc("/api/v1/events/{event_id}", middleware.Auth(handlers.GetEventDetail)).Methods("GET")
	router.HandleFunc("/api/v1/events/{event_id}", middleware.Auth(handlers.DeleteEvent)).Methods("DELETE")

	return router
}
