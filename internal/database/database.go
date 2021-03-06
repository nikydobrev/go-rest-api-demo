package database

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewDatabase - returns a pointer to a new database conneciton
func NewDatabase() (*gorm.DB, error) {
	log.Info("Setting Up Database")

	dbUserName := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbTable := os.Getenv("DB_TABLE")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=require", dbHost, dbPort, dbUserName, dbTable, dbPassword)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return db, err
	}

	return db, nil
}
