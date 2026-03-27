package middleware

import (
"github.com/go-playground/validator/v10"
"github.com/gofiber/fiber/v2"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func ValidateRequest(c *fiber.Ctx, obj interface{}) error {
	if err := c.BodyParser(obj); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
"error": "Invalid request body",
})
	}

	if err := validate.Struct(obj); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
"error": err.Error(),
		})
	}

	return nil
}
