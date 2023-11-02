package profile

import (
	"github.com/gofiber/fiber/v2"
	"implude.kr/VOAH-Backend-Core/database"
	"implude.kr/VOAH-Backend-Core/middleware"
	"implude.kr/VOAH-Backend-Core/models"
	"implude.kr/VOAH-Backend-Core/utils/validator"
)

type UpdateProfileRequest struct {
	Username    string `json:"username" validate:"required,min=1,max=30"`
	Displayname string `json:"displayname" validate:"required,min=1,max=30"`
	Position    string `json:"position" validate:"max=30"`
	Description string `json:"description" validate:"max=240"`
	DND         bool   `json:"dnd"`
}

func UpdateProfileCtrl(c *fiber.Ctx) error {
	userID, err := middleware.GetUserFromMiddleware(c)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	updateRequest := new(UpdateProfileRequest)
	if errArr := validator.ParseAndValidate(c, updateRequest); errArr != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request body",
			"errArr":  errArr,
		})
	}

	db := database.DB

	// check username is already exist
	if err := db.Where(&models.User{Username: updateRequest.Username}).First(&models.User{}).Error; err == nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Username already exist or same as before",
		})
	}

	// update user
	user := new(models.User)

	if db.First(&user, userID).Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}
	user.Username = updateRequest.Username
	user.Displayname = updateRequest.Displayname
	user.Position = updateRequest.Position
	user.Description = updateRequest.Description
	user.DND = updateRequest.DND
	if db.Save(&user).Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Profile Update Success",
	})
}
