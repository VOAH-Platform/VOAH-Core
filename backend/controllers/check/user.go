package check

import (
	"github.com/gofiber/fiber/v2"
	"implude.kr/VOAH-Backend-Core/database"
	"implude.kr/VOAH-Backend-Core/middleware"
	"implude.kr/VOAH-Backend-Core/models"
	"implude.kr/VOAH-Backend-Core/utils/permission"
)

func CheckUserCtrl(c *fiber.Ctx) error {
	userID, err := middleware.GetUserIDFromMiddleware(c)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"success": false,
		})
	}
	// find user
	db := database.DB
	foundUser := new(models.User)
	if err := db.First(&foundUser, userID).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}
	perms, err := permission.GetUserPermissionArr(foundUser)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success":    true,
		"permission": perms,
	})
}
