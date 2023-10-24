package profileimage

import (
	"encoding/base64"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"implude.kr/VOAH-Backend-Core/configs"
	"implude.kr/VOAH-Backend-Core/middleware"
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
	serverConf := configs.Env.Server
	profileName := fmt.Sprintf(serverConf.DataDir+"/user-profiles/%s.png", userUUID.String())
	_, err = os.Stat(profileName)
	if err != nil {
		return c.SendFile("./public/default-profile.webp")
	}
	return c.SendFile(profileName)
}

type UpdateImageRequest struct {
	ProfileImage string `json:"profile-image" validate:"max=409600"`
}

func UpdateImageCtrl(c *fiber.Ctx) error {
	var err error
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
	if _, err = os.Stat(imagePath); err == nil {
		// delete file
		if os.Remove(imagePath) != nil {
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
		if os.WriteFile(imagePath, decodedProfileImage, 0700) != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": "Internal server error",
			})
		}
	}

	return c.JSON(fiber.Map{
		"message": "Success",
	})
}
