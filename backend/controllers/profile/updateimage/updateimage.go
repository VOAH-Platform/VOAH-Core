package updateimage

import (
	"encoding/base64"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"implude.kr/VOAH-Backend-Core/configs"
	"implude.kr/VOAH-Backend-Core/middleware"
	"implude.kr/VOAH-Backend-Core/utils/validator"
)

type UpdateImageRequest struct {
	ProfileImage string `json:"profile-image" validate:"max=409600"`
}

func UpdateImageCtrl(c *fiber.Ctx) error {
	userID, err := middleware.GetUserID(c)
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
