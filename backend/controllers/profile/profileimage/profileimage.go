package profileimage

import (
	"encoding/base64"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"implude.kr/VOAH-Backend-Core/configs"
	"implude.kr/VOAH-Backend-Core/database"
	"implude.kr/VOAH-Backend-Core/middleware"
	"implude.kr/VOAH-Backend-Core/models"
	"implude.kr/VOAH-Backend-Core/utils/validator"
)

type GetImageRequest struct {
	UserID string `validate:"required,uuid4"`
}

func GetImageCtrl(c *fiber.Ctx) error {

	userUUID, err := uuid.Parse(c.Query("user-id", ""))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}
	db := database.DB
	user := new(models.User)

	if err := db.First(&user, userUUID).Error; err != nil && !user.Visible {
		return c.SendFile("./public/default-profile.webp")
	}
	// return profile image
	serverConf := configs.Env.Server
	return c.SendFile(fmt.Sprintf(serverConf.DataDir+"/user-profiles/%s.png", userUUID.String()))

}

type UpdateImageRequest struct {
	ProfileImage string `json:"profile-image" validate:"max=409600"`
}

func UpdateImageCtrl(c *fiber.Ctx) error {
	userID, err := middleware.GetUserIDFromMiddleware(c)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}
	updateRequest := new(UpdateImageRequest)
	if errArr := validator.ParseAndValidate(c, updateRequest); errArr != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request body",
			"errArr":  errArr,
		})
	}

	serverConf := configs.Env.Server
	imagePath := fmt.Sprintf(serverConf.DataDir+"/user-profiles/%s.png", userID.String())
	// check if file exists
	if _, err := os.Stat(imagePath); err == nil {
		// delete file
		if err := os.Remove(imagePath); err != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": "Internal server error",
			})
		}
	}
	if updateRequest.ProfileImage == "" {
		decodedProfileImage, err := base64.StdEncoding.DecodeString(updateRequest.ProfileImage)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": "Bad Profile Image",
			})
		}
		// save profile image to ./data/profiles/{uuid}.png
		err = os.WriteFile(imagePath, decodedProfileImage, 0700)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": "Internal server error",
			})
		}
	}

	return c.JSON(fiber.Map{
		"message": "Success",
	})
}
