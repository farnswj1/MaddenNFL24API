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
  opt, err := redis.ParseURL(utils.Env["REDIS_URL"])

  if err != nil {
    utils.Logger.Panic(err.Error())
  }

  cache := redis.NewClient(opt)

  if err := cache.Ping(ctx).Err(); err != nil {
    utils.Logger.Panicln(err.Error())
  }

  return cache
}
