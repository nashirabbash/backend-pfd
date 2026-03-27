package route

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/nashirabbash/backend-pfd/internal/handler"
	"github.com/nashirabbash/backend-pfd/internal/middleware"
)

func SetupRoutes(app *fiber.App) {
	authHandler := handler.NewAuthHandler()

	auth := app.Group("/api/auth")
	auth.Post("/register", authHandler.Register)
	auth.Post("/login", authHandler.Login)
	auth.Get("/me", middleware.AuthMiddleware, authHandler.Me)

	app.Get("/ws", websocket.New(handler.WebSocketHandler))

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "OK",
		})
	})
}
