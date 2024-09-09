package models

import (
  "app/utils"

  "gorm.io/driver/postgres"
  "gorm.io/gorm"
  "gorm.io/gorm/logger"
)

var DB = connectDatabase()

func connectDatabase() *gorm.DB {
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

  return database
}
