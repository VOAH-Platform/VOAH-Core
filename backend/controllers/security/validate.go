package security

import (
	"github.com/gofiber/fiber/v2"
	"github.com/pquerna/otp/totp"
	"implude.kr/VOAH-Backend-Core/database"
	"implude.kr/VOAH-Backend-Core/middleware"
	"implude.kr/VOAH-Backend-Core/utils/validator"
)

type Validate2FARequest struct {
	Code string `json:"code" validate:"required,len=6"`
}

func Validate2FACtrl(c *fiber.Ctx) error {
	user, err := middleware.GetUserFromMiddleware(c)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	} else if user.TwoFA {
		return c.Status(409).JSON(fiber.Map{
			"message": "Invalid request",
			"error":   "2FA is already enabled",
		})
	}
	validate2FARequest := new(Validate2FARequest)
	if errArr := validator.ParseAndValidate(c, validate2FARequest); errArr != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request body",
			"errArr":  errArr,
		})
	}
	if !totp.Validate(validate2FARequest.Code, user.TwoFAKey) {
		return c.Status(498).JSON(fiber.Map{
			"message": "Invalid Code",
		})
	}
	user.TwoFA = true
	db := database.DB
	if db.Save(&user).Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Success",
	})
}
