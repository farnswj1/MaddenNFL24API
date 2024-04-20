package models

import "github.com/google/uuid"

type User struct {
  ID       uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primary_key"`
  Username string    `json:"username" binding:"required,max=100" gorm:"type:varchar(100);not null"`
  Password string    `json:"password" binding:"required,max=250" gorm:"type:varchar(250);not null"`
}
