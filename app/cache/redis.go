package cache

import (
  "app/utils"
  "os"

  "github.com/redis/go-redis/v9"
)

var Cache *redis.Client

func ConnectRedis() {
  utils.Logger.Println("Connecting to Redis...")
  opt, err := redis.ParseURL(os.Getenv("REDIS_URL"))

  if err != nil {
    utils.Logger.Panic(err.Error())
  }

  Cache = redis.NewClient(opt)
  utils.Logger.Println("Connected to Redis!")
}
