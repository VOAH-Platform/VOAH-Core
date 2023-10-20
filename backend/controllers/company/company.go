package company

import (
	"github.com/gofiber/fiber/v2"
	"implude.kr/VOAH-Backend-Core/configs"
)

func GetCompanyInfo(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"success": true,
		"company": configs.Setting.Company,
	})
}
