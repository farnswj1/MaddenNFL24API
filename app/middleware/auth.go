package middleware

import (
  "app/auth"

  "github.com/gin-gonic/gin"
)

func JWTMiddleware() gin.HandlerFunc {
  return auth.JWTConfig.MiddlewareFunc()
}
