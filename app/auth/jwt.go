package auth

import (
  "app/utils"
  "time"

  "github.com/appleboy/gin-jwt/v2"
)

var JWTConfig = initJWTConfig()

func initJWTConfig() *jwt.GinJWTMiddleware {
  config, err := jwt.New(&jwt.GinJWTMiddleware{
    Key:         []byte(utils.Env["SECRET_KEY"]),
    MaxRefresh:  time.Hour,
    IdentityKey: "id",
    PayloadFunc: payloadFunc,
    IdentityHandler: identityHandler,
    Authenticator: authenticator,
    Authorizator: authorizator,
    Unauthorized: unauthorized,
  })

  if err != nil {
    utils.Logger.Panic(err.Error())
  }

  return config
}
