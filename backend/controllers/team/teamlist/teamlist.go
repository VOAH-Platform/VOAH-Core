package teamlist

import (
	"github.com/gofiber/fiber/v2"
	"implude.kr/VOAH-Backend-Core/database"
	"implude.kr/VOAH-Backend-Core/models"
)

func TeamListCtrl(c *fiber.Ctx) error {
	db := database.DB

	// fetch public teams
	var teamList []models.Team
	if db.Where(&models.Team{Visible: true}).Find(&teamList).Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Success",
		"teams":   teamList,
	})
}
