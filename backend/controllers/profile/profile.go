package profile

import (
	"context"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"implude.kr/VOAH-Backend-Core/database"
	"implude.kr/VOAH-Backend-Core/middleware"
	"implude.kr/VOAH-Backend-Core/models"
	"implude.kr/VOAH-Backend-Core/utils/validator"
)

type GetProfileRequest struct {
	UserID string `validate:"required,uuid4"`
}

func GetProfileCtrl(c *fiber.Ctx) error {
	_, err := middleware.GetUserFromMiddleware(c)

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	profileRequest := GetProfileRequest{
		UserID: c.Query("user-id"),
	}

	if errArr := validator.VOAHValidator.Validate(profileRequest); len(errArr) != 0 {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request",
			"error":   errArr,
		})
	}

	db := database.DB
	user := new(models.User)

	if db.First(&user, uuid.MustParse(profileRequest.UserID)).Error != nil && !user.Visible {
		return c.Status(400).JSON(fiber.Map{
			"message": "User not found",
		})
	}

	// get last activity and refresh
	var lastActivity int64
	var lastRefresh int64
	ctx := context.Background()

	lastActivityRedis := database.Redis.LastActivityRedis
	lastActivityStr, err := lastActivityRedis.Get(ctx, profileRequest.UserID).Result()
	if err != nil {
		lastActivity = time.Now().Unix()
	} else {
		lastActivity, err = strconv.ParseInt(lastActivityStr, 10, 64)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": "Internal server error",
			})
		}
	}

	lastRefreshRedis := database.Redis.LastRefreshRedis
	lastRefreshStr, err := lastRefreshRedis.Get(ctx, profileRequest.UserID).Result()
	if err != nil {
		lastRefresh = time.Now().Unix()
	} else {
		lastRefresh, err = strconv.ParseInt(lastRefreshStr, 10, 64)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": "Internal server error",
			})
		}
	}

	return c.JSON(fiber.Map{
		"message":       "Success",
		"user":          user,
		"last-activity": lastActivity,
		"last-refresh":  lastRefresh,
	})
}

type UpdateProfileRequest struct {
	Username    string `json:"username" validate:"required,min=1,max=30"`
	Displayname string `json:"displayname" validate:"required,min=1,max=30"`
	Position    string `json:"position" validate:"max=30"`
	Description string `json:"description" validate:"max=240"`
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
	if db.Save(&user).Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Profile Update Success",
	})
}
