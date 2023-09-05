package passreset

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"implude.kr/VOAH-Backend-Core/database"
	"implude.kr/VOAH-Backend-Core/models"
	"implude.kr/VOAH-Backend-Core/utils/validator"
)

type CheckPassResetRequest struct {
	Email string `json:"email" validate:"required,email"`
	Code  string `json:"code" validate:"required,uuid4"`
}

func CheckPassResetCtrl(c *fiber.Ctx) error {
	checkPassResetRequest := new(CheckPassResetRequest)
	if err := c.BodyParser(checkPassResetRequest); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}
	if errArr := validator.VOAHValidator.Validate(checkPassResetRequest); len(errArr) != 0 {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request",
			"error":   errArr,
		})
	}

	// check if unvalidated user exists
	ctx := context.Background()
	redis := database.Redis.PasswordResetRedis
	if redis.Exists(ctx, checkPassResetRequest.Code).Val() == 0 {
		return c.Status(401).JSON(fiber.Map{
			"message": "Invalid code",
		})
	} else if redis.Get(ctx, checkPassResetRequest.Code).Val() != checkPassResetRequest.Email {
		return c.Status(401).JSON(fiber.Map{
			"message": "Invalid code",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Valid code",
	})
}

type SubmitPassResetRequest struct {
	Code     string `json:"code" validate:"required,uuid4"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=32"`
}

func SubmitPassResetCtrl(c *fiber.Ctx) error {
	passResetRequest := new(SubmitPassResetRequest)
	if err := c.BodyParser(passResetRequest); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	if errArr := validator.VOAHValidator.Validate(passResetRequest); len(errArr) != 0 {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request",
			"error":   errArr,
		})
	}

	// check if pass reset code exists
	ctx := context.Background()
	passResetRedis := database.Redis.PasswordResetRedis
	if passResetRedis.Exists(ctx, passResetRequest.Code).Val() == 0 {
		return c.Status(401).JSON(fiber.Map{
			"message": "Invalid code",
		})
	} else if passResetRedis.Get(ctx, passResetRequest.Code).Val() != passResetRequest.Email {
		return c.Status(401).JSON(fiber.Map{
			"message": "Invalid code",
		})
	}
	passResetRedis.Del(ctx, passResetRequest.Code)

	// update password
	db := database.DB
	user := new(models.User)
	if err := db.Where(&models.User{Email: passResetRequest.Email}).First(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	pwHash, err := bcrypt.GenerateFromPassword([]byte(passResetRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}
	user.PWHash = string(pwHash)
	if err := db.Save(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Pass Reset Success",
	})
}
