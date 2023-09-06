package info

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"implude.kr/VOAH-Backend-Core/database"
	"implude.kr/VOAH-Backend-Core/middleware"
	"implude.kr/VOAH-Backend-Core/models"
)

func GetModuleList(c *fiber.Ctx) error {
	_, err := middleware.GetUserIDFromMiddleware(c)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"success": false,
		})
	}

	// get enabled modules
	db := database.DB

	var modules []models.Module

	if err := db.Where(&models.Module{Enabled: true}).Find(&modules).Error; err != nil && err != gorm.ErrRecordNotFound {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"modules": modules,
	})
}
