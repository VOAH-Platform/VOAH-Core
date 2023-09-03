package getimage

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"implude.kr/VOAH-Backend-Core/configs"
	"implude.kr/VOAH-Backend-Core/database"
	"implude.kr/VOAH-Backend-Core/middleware"
	"implude.kr/VOAH-Backend-Core/models"
)

type GetImageRequest struct {
	UserID string `validate:"required,uuid4"`
}

func GetImageCtrl(c *fiber.Ctx) error {
	_, err := middleware.GetUserID(c)
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
