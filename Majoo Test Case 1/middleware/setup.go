package middleware

import (
	"github.com/gofiber/fiber/v2"
	"majoo.id/satu/models"
)

// Set context default value
func Setup(c *fiber.Ctx) error {
	c.Locals("ctx", &models.Context{})
	return c.Next()
}
