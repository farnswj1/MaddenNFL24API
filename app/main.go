package main

import (
	"app/controllers"
	"app/models"
	"app/utils"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func getRouter() *gin.Engine {
	gin.SetMode(utils.GetenvOrDefault("GIN_MODE", "debug"))
	r := gin.Default()
	r.GET("/", controllers.Root)
	r.GET("/players", controllers.FindPlayers)
	return r
}

func main() {
	godotenv.Load()
	models.ConnectDatabase()
	r := getRouter()
	r.Run()
}
