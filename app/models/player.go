package models

import "app/serializers"

type Player struct {
  ID uint32 `json:"id" gorm:"primary_key;autoIncrement"`
  serializers.Player
}
