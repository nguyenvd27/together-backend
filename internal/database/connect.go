package database

import (
	"log"
	"os"
	"fmt"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Use PostgreSql
func ConnectDB() *gorm.DB {
	skipLoadEvn := os.Getenv("SKIP_LOAD_ENV")
	if skipLoadEvn == "" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("Some error occured. Err: %s", err)
			return &gorm.DB{}
		}
	}

	userName := os.Getenv("USER_NAME")
	password := os.Getenv("PASSWORD")
	host := os.Getenv("DATABASE_HOST")
	port := os.Getenv("DATABASE_PORT")
	databaseName := os.Getenv("DATABASE_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", host, userName, password, databaseName, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	return db
}
