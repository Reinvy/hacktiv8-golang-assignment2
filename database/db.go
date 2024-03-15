package database

import (
	"assignment2/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

// InitDB initializes the database connection
func InitDB() {
	// Build the connection string
	connectionString := "host=localhost port=5432 user=postgres password=root dbname=assignment2 sslmode=disable"

	// Connect to the database
	var err error
	db, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Migrate the database schema
	db.AutoMigrate(&models.Order{}, &models.Item{})
}

// GetDB returns the database instance
func GetDB() *gorm.DB {
	return db
}
