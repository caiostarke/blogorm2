package database

import (
	"log"

	"github.com/caiostarke/studyToAdvanced/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func Connect() {
	DB, err = gorm.Open(mysql.Open("root:eduardo3@tcp(localhost:3306)/blogorm"))

	if err != nil {
		log.Fatal("Failed connect to DB")
	}
}

func AutoMigrate() {
	DB.AutoMigrate(&models.Post{}, &models.User{}, &models.Tag{})
}
