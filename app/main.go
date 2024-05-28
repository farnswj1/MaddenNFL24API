package main

import (
  "app/auth"
  "app/controllers"
  "app/middleware"
  "app/utils"

  "github.com/gin-gonic/gin"
)

func getRouter() *gin.Engine {
  router := gin.Default()
  router.RemoveExtraSlash = true
  router.Use(middleware.CorsMiddleware())
  router.LoadHTMLGlob("templates/*")
  router.NoRoute(controllers.PageNotFound)
  router.NoMethod(controllers.MethodNotAllowed)
  router.GET("/", controllers.Root)

  authGroup := router.Group("/auth")
  authGroup.POST("/login", auth.JWTConfig.LoginHandler)
  authGroup.POST("/refresh", auth.JWTConfig.RefreshHandler)

  playersGroup := router.Group("/players", middleware.RateLimiterMiddleware("app", 10, 60))
  playersGroup.GET("/", controllers.FindPlayers)
  playersGroup.GET("/:id", controllers.FindPlayer)

  modifyPlayersGroup := playersGroup.Group("", middleware.JWTMiddleware())
  modifyPlayersGroup.POST("/", controllers.CreatePlayer)
  modifyPlayersGroup.PUT("/:id", controllers.UpdatePlayer)
  modifyPlayersGroup.DELETE("/:id", controllers.DeletePlayer)

  return router
}

func init() {
  gin.SetMode(utils.Env.Get("GIN_MODE", "debug"))
}

func main() {
  router := getRouter()
  router.Run()
}
