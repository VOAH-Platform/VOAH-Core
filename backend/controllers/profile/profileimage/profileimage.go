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
	_, err := middleware.GetUserIDFromMiddleware(c)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	profileRequest := GetImageRequest{
		UserID: c.Query("user-id"),
	}
	db := database.DB
	user := new(models.User)

	if err := db.First(&user, uuid.MustParse(profileRequest.UserID)).Error; err != nil && !user.Visible {
		return c.Status(400).JSON(fiber.Map{
			"message": "User not found",
		})
	}
	// return profile image
	serverConf := configs.Env.Server
	return c.SendFile(fmt.Sprintf(serverConf.DataDir+"/user-profiles/%s.png", profileRequest.UserID))

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
