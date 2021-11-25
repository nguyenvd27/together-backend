package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"together-backend/internal/router"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func initRouter(port string) {
	router := router.New()

	corsWrapper := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "DELETE", "PUT"},
		AllowedHeaders: []string{"Content-Type", "Authorization", "Origin", "Accept", "*"},
	})

	if port == "" {
		port = "8001" // Default port if not specified
	}
	log.Fatal(http.ListenAndServe(":"+port, corsWrapper.Handler(router)))
}

func main() {
	fmt.Println("Together Backend App.....")

	skipLoadEvn := os.Getenv("SKIP_LOAD_ENV")
	if skipLoadEvn == "" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("Some error occured. Err: %s", err)
			return
		}
	}
	port := os.Getenv("PORT")

	initRouter(port)
}
