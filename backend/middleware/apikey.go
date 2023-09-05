package middleware

import (
	"github.com/gofiber/fiber/v2"
	"implude.kr/VOAH-Backend-Core/configs"
)

func CheckAPIKeyMiddleware(c *fiber.Ctx) error {
	apiKey := c.Get("API-KEY", "")
	if apiKey != configs.APIKey {
		return c.Status(403).JSON(fiber.Map{
			"message": "Invalid API Key",
		})
	}
	return c.Next()
}
