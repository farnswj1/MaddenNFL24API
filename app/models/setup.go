package models

import (
	"app/utils"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	utils.Logger.Println("Connecting to database...")

	database, err := gorm.Open(
		postgres.Open(os.Getenv("DATABASE_URL")),
		&gorm.Config{},
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
