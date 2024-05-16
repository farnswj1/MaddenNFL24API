package models

import (
  "app/utils"

  "gorm.io/driver/postgres"
  "gorm.io/gorm"
  "gorm.io/gorm/logger"
)

var DB *gorm.DB

func ConnectDatabase() {
  if DB != nil {
    return
  }

  utils.Logger.Println("Connecting to database...")

  config := &gorm.Config{}
  mode := utils.Env.Get("GIN_MODE", "debug")

  if mode == "release" {
    config.Logger = logger.Default.LogMode(logger.Silent)
  }

  database, err := gorm.Open(
    postgres.Open(utils.Env["DATABASE_URL"]),
    config,
  )

  if err != nil {
    utils.Logger.Panic(err.Error())
  }

  if err = database.AutoMigrate(&Player{}); err != nil {
    utils.Logger.Panic(err.Error())
  }

  if err = database.AutoMigrate(&User{}); err != nil {
    utils.Logger.Panic(err.Error())
  }

  DB = database
  utils.Logger.Println("Connected to database!")
}

func init() {
  ConnectDatabase()
}
