package database

import (
	"fmt"
	"gorm.io/gorm"
	m "last-level/models"
	
)

func MigrateDb(db *gorm.DB) error {
	err := db.AutoMigrate(
		&m.Admin{},
		
		
	)

	

	if err != nil {
		return fmt.Errorf("error during migration: %w", err)
	}

	fmt.Println("Database migration completed successfully")
	
	return nil
}