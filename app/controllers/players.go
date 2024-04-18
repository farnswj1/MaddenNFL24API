package controllers

import (
  "app/models"
  "net/http"

  "github.com/gin-gonic/gin"
)

func Root(c *gin.Context) {
  c.JSON(http.StatusOK, "This is a test!!!!")
}

func FindPlayers(c *gin.Context) {
  var players []models.Player
  models.DB.Find(&players)
  c.JSON(http.StatusOK, players)
}

func CreatePlayer(c *gin.Context) {
  var player models.Player

  if err := c.ShouldBindJSON(&player); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  models.DB.Create(&player)
  c.JSON(http.StatusCreated, player)
}

func FindPlayer(c *gin.Context) {
  id := c.Param("id")
  var player models.Player

  if err := models.DB.First(&player, "id = ?", id).Error; err != nil {
    c.JSON(http.StatusNotFound, gin.H{"error": "Player not found"})
    return
  }

  c.JSON(http.StatusOK, player)
}
