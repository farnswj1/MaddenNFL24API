package controllers

import (
  "app/models"
  "app/utils"
  "net/http"

  "github.com/gin-gonic/gin"
)

func Root(c *gin.Context) {
  var count int64
  models.DB.Model(&models.Player{}).Count(&count)
  numPlayers := utils.Formatter.Sprintf("%d", count)
  c.HTML(http.StatusOK, "index.html", gin.H{"numPlayers": numPlayers})
}
