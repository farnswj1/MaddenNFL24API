package controllers

import (
  "app/models"
  "net/http"

  "github.com/gin-gonic/gin"
)

func FindPlayers(c *gin.Context) {
  var players []models.Player
  models.DB.Find(&players)
  c.JSON(http.StatusOK, players)
}

func CreatePlayer(c *gin.Context) {
  var player models.Player

  if err := c.ShouldBindJSON(&player); err != nil {
    c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  if err := models.DB.Create(&player).Error; err != nil {
    c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  c.JSON(http.StatusCreated, player)
}

func FindPlayer(c *gin.Context) {
  id := c.Param("id")
  var player models.Player

  if err := models.DB.First(&player, "id = ?", id).Error; err != nil {
    c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
    return
  }

  c.JSON(http.StatusOK, player)
}

func UpdatePlayer(c *gin.Context) {
  id := c.Param("id")
  var player models.PlayerWithID

  if err := c.ShouldBindJSON(&player); err != nil {
    c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  if id != player.ID.String() {
    c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "ID in request body does not match URL"})
    return
  }

  result := models.DB.Model(&models.Player{}).Where("id = ?", id).Updates(&player)

  if err := result.Error; err != nil {
    c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
    return
  }

  if result.RowsAffected == 0 {
    c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Player not found"})
    return
  }

  c.JSON(http.StatusOK, player)
}

func DeletePlayer(c *gin.Context) {
  id := c.Param("id")
  var player models.Player
  result := models.DB.Delete(&player, "id = ?", id)

  if err := result.Error; err != nil {
    c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
    return
  }

  if result.RowsAffected == 0 {
    c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "Player not found"})
    return
  }

  c.JSON(http.StatusNoContent, nil)
}
