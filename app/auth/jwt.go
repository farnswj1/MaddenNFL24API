package auth

import (
  "app/utils"
  "os"
  "time"

  "github.com/appleboy/gin-jwt/v2"
)

var JWTConfig = func() *jwt.GinJWTMiddleware {
  config, err := jwt.New(&jwt.GinJWTMiddleware{
    Key:         []byte(os.Getenv("SECRET_KEY")),
    MaxRefresh:  time.Hour,
    IdentityKey: "id",
    PayloadFunc: PayloadFunc,
    IdentityHandler: IdentityHandler,
    Authenticator: Authenticator,
    Authorizator: Authorizator,
    Unauthorized: Unauthorized,
  })

  if err != nil {
    utils.Logger.Panic(err.Error())
  }

  return config
}()
