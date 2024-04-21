package main

import (
  "app/auth"
  "app/cache"
  "app/controllers"
  "app/middleware"
  "app/models"
  "app/utils"

  "github.com/gin-gonic/gin"
  "github.com/joho/godotenv"
)

func getRouter() *gin.Engine {
  gin.SetMode(utils.GetenvOrDefault("GIN_MODE", "debug"))
  router := gin.Default()
  router.RemoveExtraSlash = true
  router.Use(middleware.RateLimiter("app", 10, 60))
  router.NoRoute(controllers.PageNotFound)
  router.NoMethod(controllers.MethodNotAllowed)
  router.GET("/", controllers.Root)

  authGroup := router.Group("/auth")
  authGroup.POST("/login", auth.JWTConfig.LoginHandler)
  authGroup.POST("/refresh", auth.JWTConfig.RefreshHandler)

  playersGroup := router.Group("/players")
  playersGroup.GET("/", controllers.FindPlayers)
  playersGroup.GET("/:id", controllers.FindPlayer)

  modifyPlayersGroup := playersGroup.Group("", middleware.JWTMiddleware())
  modifyPlayersGroup.POST("/", controllers.CreatePlayer)
  modifyPlayersGroup.PUT("/:id", controllers.UpdatePlayer)
  modifyPlayersGroup.DELETE("/:id", controllers.DeletePlayer)

  return router
}

func main() {
  godotenv.Load()
  models.ConnectDatabase()
  cache.ConnectRedis()
  router := getRouter()
  router.Run()
}
