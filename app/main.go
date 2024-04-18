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
  router := gin.Default()
  router.RemoveExtraSlash = true
  router.NoRoute(controllers.PageNotFound)
  router.NoMethod(controllers.MethodNotAllowed)
  router.GET("/", controllers.Root)

  group := router.Group("/players")
  group.GET("/", controllers.FindPlayers)
  group.POST("/", controllers.CreatePlayer)
  group.GET("/:id", controllers.FindPlayer)

  return router
}

func main() {
  godotenv.Load()
  models.ConnectDatabase()
  router := getRouter()
  router.Run()
}
