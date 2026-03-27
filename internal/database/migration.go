package database

import (
	"fmt"
	"log"

	"github.com/nashirabbash/backend-pfd/internal/model"
)

func AutoMigrate() error {
	if DB == nil {
		return fmt.Errorf("database is not initialized")
	}

	err := DB.AutoMigrate(
		&model.User{},
	)

	if err != nil {
		log.Printf("Auto migration failed: %v", err)
		return err
	}

	log.Println("✓ Database migration completed successfully")
	return nil
}
