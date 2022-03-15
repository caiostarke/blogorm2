package database

import (
	"fmt"
	"log"
	"os"

	"github.com/caiostarke/studyToAdvanced/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Connect() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	MYSQL_ROOT_PASSWORD := os.Getenv("MYSQL_ROOT_PASSWORD")
	DB_USER := os.Getenv("DB_USER")
	DB_HOST := os.Getenv("DB_HOST")
	DB_NAME := os.Getenv("DB_NAME")

	URL := fmt.Sprintf("%s:%s@tcp(%s)/%s", DB_USER, MYSQL_ROOT_PASSWORD, DB_HOST, DB_NAME)

	DB, err = gorm.Open(mysql.Open(URL))

	if err != nil {
		log.Fatal("Failed connect to DB")
	}
}

func AutoMigrate() {
	DB.AutoMigrate(&models.Post{}, &models.User{}, &models.Tag{})
}
