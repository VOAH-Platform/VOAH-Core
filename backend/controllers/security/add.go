package security

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"implude.kr/VOAH-Backend-Core/configs"
	"implude.kr/VOAH-Backend-Core/database"
	"implude.kr/VOAH-Backend-Core/middleware"
	"implude.kr/VOAH-Backend-Core/utils/logger"
	"implude.kr/VOAH-Backend-Core/utils/twofa"
)

func Add2FACtrl(c *fiber.Ctx) error {
	user, err := middleware.GetUserFromMiddleware(c)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	if user.TwoFA {
		return c.Status(409).JSON(fiber.Map{
			"message": "Invalid request",
			"error":   "2FA is already enabled",
		})
	}
	tfa := twofa.TwoFA{
		Issuer: configs.Setting.Company.Name,
		Email:  user.Email,
	}
	if tfa.NewKey() != nil {
		log := logger.Logger
		log.Error(fmt.Sprintf("Error while generating new key for user %s", user.Email))
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	user.TwoFAKey = tfa.Key.Secret()
	db := database.DB
	if db.Save(&user).Error != nil {
		log := logger.Logger
		log.Error(fmt.Sprintf("Error while saving new key for user %s", user.Email))
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}
	return c.JSON(fiber.Map{
		"message":   "Success",
		"key":       tfa.Key.Secret(),
		"qr-base64": tfa.ImageBase64,
	})
}
