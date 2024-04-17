package models

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(
		postgres.Open(os.Getenv("DATABASE_URL")),
		&gorm.Config{},
	)

	if err != nil {
		panic("Failed to connect to database!")
	}

	if err = database.AutoMigrate(&Player{}); err != nil {
		panic("Faield to make migrations to database!")
	}

	DB = database
}
