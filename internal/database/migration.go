package database

import (
"log"

"github.com/nashirabbash/backend-pfd/internal/model"
)

func AutoMigrate() error {
	if DB == nil {
		return nil
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
