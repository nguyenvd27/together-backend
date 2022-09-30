package database

// import (
// 	"log"
// 	"os"

// 	"github.com/joho/godotenv"
// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// )

// Use mysql
// func ConnectDB() *gorm.DB {
// 	skipLoadEvn := os.Getenv("SKIP_LOAD_ENV")
// 	if skipLoadEvn == "" {
// 		err := godotenv.Load(".env")
// 		if err != nil {
// 			log.Fatalf("Some error occured. Err: %s", err)
// 			return &gorm.DB{}
// 		}
// 	}

// 	userName := os.Getenv("USER_NAME")
// 	password := os.Getenv("PASSWORD")
// 	databasePath := os.Getenv("DATABASE_PATH")
// 	databaseName := os.Getenv("DATABASE_NAME")

// 	dsn := userName + ":" + password + "@" + databasePath + databaseName + "?charset=utf8mb4&parseTime=True&loc=Local"

// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	return db
// }
