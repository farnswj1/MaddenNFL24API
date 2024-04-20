package auth

import (
  "app/models"

  "github.com/appleboy/gin-jwt/v2"
  "github.com/gin-gonic/gin"
  "github.com/google/uuid"
  "golang.org/x/crypto/bcrypt"
)

func PayloadFunc(data interface{}) jwt.MapClaims {
  if user, ok := data.(*models.User); ok {
    return jwt.MapClaims{
      "id": user.ID.String(),
    }
  }
  return jwt.MapClaims{}
}

func IdentityHandler(c *gin.Context) interface{} {
  claims := jwt.ExtractClaims(c)
  id, _ := uuid.Parse(claims["id"].(string))
  var user models.User

  if err := models.DB.First(&user, "id = ?", id).Error; err != nil {
    return nil
  }

  return user
}

func Authenticator(c *gin.Context) (interface{}, error) {
  var credentials, user models.User

  if err := c.ShouldBind(&credentials); err != nil {
    return "", jwt.ErrMissingLoginValues
  }

  if err := models.DB.Where("username = ?", credentials.Username).First(&user).Error; err != nil {
    return nil, jwt.ErrFailedAuthentication
  }

  if err := bcrypt.CompareHashAndPassword(
    []byte(user.Password),
    []byte(credentials.Password),
  ); err != nil {
    return nil, jwt.ErrFailedAuthentication
  }

  return &user, nil
}

func Authorizator(data interface{}, c *gin.Context) bool {
  return true
}

func Unauthorized(c *gin.Context, code int, message string) {
  c.AbortWithStatusJSON(code, gin.H{"error": message})
}
