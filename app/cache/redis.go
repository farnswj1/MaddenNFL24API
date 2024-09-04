package cache

import (
  "app/utils"
  "context"

  "github.com/redis/go-redis/v9"
)

var (
  Cache = connectRedis()
  ctx = context.Background()
)

func connectRedis() *redis.Client {
  utils.Logger.Println("Connecting to Redis...")
  opt, err := redis.ParseURL(utils.Env["REDIS_URL"])

  if err != nil {
    utils.Logger.Panic(err.Error())
  }

  cache := redis.NewClient(opt)

  if err := cache.Ping(ctx).Err(); err != nil {
    utils.Logger.Panicln(err.Error())
  }

  utils.Logger.Println("Connected to Redis!")
  return cache
}
