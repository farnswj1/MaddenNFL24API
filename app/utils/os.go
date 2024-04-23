package utils

import (
  "os"
  "strings"

  "github.com/joho/godotenv"
)

type EnvMap map[string]string
var Env EnvMap

func LoadEnv() {
  godotenv.Load()
  items := EnvMap{}

  for _, item := range os.Environ() {
    splits := strings.Split(item, "=")
    key := splits[0]
    val := splits[1]
    items[key] = val
  }

  Env = items
}

func GetenvOrDefault(key, fallback string) string {
  if value, exists := Env[key]; exists {
    return value
  }

  return fallback
}
