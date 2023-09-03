package updateprofile

import (
	"github.com/gofiber/fiber/v2"
	"implude.kr/VOAH-Backend-Core/database"
	"implude.kr/VOAH-Backend-Core/middleware"
	"implude.kr/VOAH-Backend-Core/models"
	"implude.kr/VOAH-Backend-Core/utils/validator"
)

type UpdateProfileRequest struct {
	Displayname string `json:"displayname" validate:"required,min=1,max=30"`
	Position    string `json:"position" validate:"max=30"`
}

func UpdateProfileCtrl(c *fiber.Ctx) error {
	userID, err := middleware.GetUserID(c)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}
	updateRequest := new(UpdateProfileRequest)
	if err := c.BodyParser(updateRequest); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}
	if errArr := validator.VOAHValidator.Validate(updateRequest); len(errArr) != 0 {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request",
			"error":   errArr,
		})
	}
	db := database.DB

	// update user
	user := new(models.User)

	if err := db.First(&user, userID).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}
	user.Displayname = updateRequest.Displayname
	user.Position = updateRequest.Position
	if err := db.Save(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Profile Update Success",
	})
}
