package db

import (
	"go-virtual-currency/models"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error
	DB, err = gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to SQLite:", err)
	}
	err = DB.AutoMigrate(&models.Paste{})
	if err != nil {
		log.Fatal("Migration failed:", err)
	}
}
