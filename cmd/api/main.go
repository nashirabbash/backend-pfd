package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/nashirabbash/backend-pfd/internal/config"
	"github.com/nashirabbash/backend-pfd/internal/database"
	"github.com/nashirabbash/backend-pfd/internal/route"
)

func main() {
	log.Println("📋 Loading configuration...")
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	log.Println("✓ Configuration loaded successfully")

	log.Println("🗄️  Initializing database connection...")
	if err := database.InitDB(cfg); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer database.CloseDB()

	log.Println("🔄 Running database migrations...")
	if err := database.AutoMigrate(); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	log.Println("🚀 Creating Fiber application...")
	app := fiber.New(fiber.Config{
		AppName: "PFD Backend API v1.0.0",
	})

	log.Println("⚙️  Registering middleware...")
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE,PATCH,OPTIONS",
		AllowHeaders: "Content-Type,Authorization",
	}))

	log.Println("🛣️  Setting up routes...")
	route.SetupRoutes(app)

	log.Printf("🌐 Server starting on port %s...\n", cfg.Port)
	if err := app.Listen(":" + cfg.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
