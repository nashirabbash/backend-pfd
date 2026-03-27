package database

import (
	"fmt"
	"log"

	"github.com/nashirabbash/backend-pfd/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(cfg *config.Config) error {
	dsn := cfg.GetDSN()

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		return fmt.Errorf("database connection failed: %w", err)
	}

	log.Println("✓ Database connected successfully")
	DB = db
	return nil
}

func GetDB() *gorm.DB {
	return DB
}

func CloseDB() error {
	if DB == nil {
		return nil
	}
	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
