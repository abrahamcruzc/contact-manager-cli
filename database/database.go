package database

import (
	"log"

	"github.com/abrahamcruzc/contact-manager-cli/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	var err error
	DB, err = gorm.Open(sqlite.Open("contacts.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
		return err
	}
	
	if err = DB.AutoMigrate(&models.Contact{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
		return err
	}
	
	log.Println("Database connected")
	return nil
}

