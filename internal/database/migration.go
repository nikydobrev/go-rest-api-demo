package database

import (
	"github.com/nikydobrev/go-rest-api-demo/internal/comment"
	"gorm.io/gorm"
)

// MigrateDB - migrates our database and creates our comment table
func MigrateDB(db *gorm.DB) error {
	if err := db.AutoMigrate(&comment.Comment{}); err != nil {
		return err
	}
	return nil
}
