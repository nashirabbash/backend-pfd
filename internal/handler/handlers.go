package handler

import (
	"fmt"
	"log"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/nashirabbash/backend-pfd/internal/database"
	"github.com/nashirabbash/backend-pfd/internal/dto"
	"github.com/nashirabbash/backend-pfd/internal/middleware"
	"github.com/nashirabbash/backend-pfd/internal/repository"
	"github.com/nashirabbash/backend-pfd/internal/service"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler() *AuthHandler {
	userRepo := repository.NewUserRepository(database.GetDB())
	authService := service.NewAuthService(userRepo)
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var req dto.RegisterRequest
	if err := middleware.ValidateRequest(c, &req); err != nil {
		return err
	}

	resp, err := h.authService.Register(&req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(resp)
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req dto.LoginRequest
	if err := middleware.ValidateRequest(c, &req); err != nil {
		return err
	}

	resp, err := h.authService.Login(&req)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(resp)
}

func (h *AuthHandler) Me(c *fiber.Ctx) error {
	userID := c.Locals("user_id")
	email := c.Locals("email")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"user_id": userID,
		"email":   email,
	})
}

func WebSocketHandler(c *websocket.Conn) {
	log.Println("✓ WebSocket client connected")

	for {
		messageType, message, err := c.ReadMessage()
		if err != nil {
			log.Printf("WebSocket read error: %v", err)
			break
		}

		log.Printf("Received message (type %d): %s", messageType, message)

		echoMessage := fmt.Sprintf("Echo: %s", string(message))
		err = c.WriteMessage(messageType, []byte(echoMessage))
		if err != nil {
			log.Printf("WebSocket write error: %v", err)
			break
		}
	}

	log.Println("✗ WebSocket client disconnected")
}
