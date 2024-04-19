package controllers

import (
  "net/http"

  "github.com/gin-gonic/gin"
)

func Root(c *gin.Context) {
  c.JSON(http.StatusOK, "Welcome to Madden NFL 24 API!")
}
