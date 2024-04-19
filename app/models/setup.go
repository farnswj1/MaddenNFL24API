package models

import (
  "app/utils"
  "os"

  "gorm.io/driver/postgres"
  "gorm.io/gorm"
  "gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDatabase() {
  utils.Logger.Println("Connecting to database...")

  config := &gorm.Config{}
  mode := utils.GetenvOrDefault("GIN_MODE", "debug")

  if mode == "release" {
    config.Logger = logger.Default.LogMode(logger.Silent)
  }

  database, err := gorm.Open(
    postgres.Open(os.Getenv("DATABASE_URL")),
    config,
  )

  if err != nil {
    utils.Logger.Panic(err.Error())
  }

  if err = database.AutoMigrate(&Player{}); err != nil {
    utils.Logger.Panic(err.Error())
  }

  DB = database
  utils.Logger.Println("Connected to database!")
}
