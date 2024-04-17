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
	c.JSON(http.StatusOK, gin.H{"data": players})
}
