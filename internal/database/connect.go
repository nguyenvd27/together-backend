package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const USER_NAME = "root"
const PASSWORD = "12345678"
const DATABASE_PATH = "tcp(127.0.0.1:3306)/test_pbl"

func ConnectDB() *gorm.DB {
	dsn := USER_NAME + ":" + PASSWORD + "@" + DATABASE_PATH + "?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	return db
}
