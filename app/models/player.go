package models

import (
	"app/serializers"

	"github.com/google/uuid"
)

type Player struct {
  ID uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primary_key;"`
  serializers.Player
}
