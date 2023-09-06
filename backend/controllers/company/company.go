package company

import (
	"github.com/gofiber/fiber/v2"
	"implude.kr/VOAH-Backend-Core/configs"
	"implude.kr/VOAH-Backend-Core/middleware"
)

func GetCompanyInfo(c *fiber.Ctx) error {
	_, err := middleware.GetUserIDFromMiddleware(c)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"success": false,
		})
	}
	return c.JSON(fiber.Map{
		"success": true,
		"company": configs.Setting.Company,
	})
}
