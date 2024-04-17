package utils

import "os"

func GetenvOrDefault(key, fallback string) string {
  if value, exists := os.LookupEnv(key); exists {
    return value
  }

  return fallback
}
