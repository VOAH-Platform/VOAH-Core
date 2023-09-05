package check

import (
	"github.com/gofiber/fiber/v2"
	"implude.kr/VOAH-Backend-Core/middleware"
)

func GetPermissionCtrl(c *fiber.Ctx) error {
	_, err := middleware.GetUserIDFromMiddleware(c)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"success":    false,
			"permission": []string{},
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"success":    true,
		"permission": []string{},
	})
}
