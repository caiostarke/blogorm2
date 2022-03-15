package main

import (
	"github.com/caiostarke/studyToAdvanced/database"
	"github.com/caiostarke/studyToAdvanced/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	database.Connect()
	database.AutoMigrate()

	router.Static("/assets", "./assets")
	routes.Setup(router)

	router.Run()
}
