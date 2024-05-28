package middleware

import (
  "app/cache"
  "app/utils"
  "net/http"
  "time"

  "github.com/gin-gonic/gin"
)

func RateLimiterMiddleware(endpoint string, limit, duration int64) gin.HandlerFunc {
  return func(c *gin.Context) {
    context := c.Request.Context()
    key := endpoint + ":" + c.ClientIP()
    numRequests, err := cache.Cache.Incr(context, key).Result()

    if err != nil {
      utils.Logger.Panic(err.Error())
    }

    if numRequests == 1 {
      expiration := time.Duration(duration) * time.Second
      cache.Cache.Expire(context, key, expiration)
    }

    if numRequests > limit {
      c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "Too many requests"})
      return
    }

    c.Next()
  }
}
