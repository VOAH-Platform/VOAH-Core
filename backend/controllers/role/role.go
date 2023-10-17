package role

import (
	"github.com/gofiber/fiber/v2"
	"implude.kr/VOAH-Backend-Core/middleware"
	"implude.kr/VOAH-Backend-Core/utils/checkperm"
)

func GetMyRoles(c *fiber.Ctx) error {
	user, err := middleware.GetUserFromMiddleware(c)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"success": false,
		})
	}
	userRoles, err := checkperm.GetUserRoleArr(user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"roles":   userRoles,
	})
}
