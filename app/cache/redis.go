package cache

import (
  "app/utils"

  "github.com/redis/go-redis/v9"
)

var Cache *redis.Client

func connectRedis() {
  if Cache != nil {
    return
  }

  utils.Logger.Println("Connecting to Redis...")
  opt, err := redis.ParseURL(utils.Env["REDIS_URL"])

  if err != nil {
    utils.Logger.Panic(err.Error())
  }

  Cache = redis.NewClient(opt)
  utils.Logger.Println("Connected to Redis!")
}

func init() {
  connectRedis()
}
