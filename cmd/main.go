package main

import (
	"fmt"
	"log"
	"net/http"
	"together-backend/internal/router"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func initRouter() {
	router := router.New()

	corsWrapper := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "DELETE", "PUT"},
		AllowedHeaders: []string{"Content-Type", "Authorization", "Origin", "Accept", "*"},
	})

	log.Fatal(http.ListenAndServe(":8001", corsWrapper.Handler(router)))
}

func main() {
	fmt.Println("Together Backend App")

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
		return
	}

	initRouter()
}
