package profile

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"implude.kr/VOAH-Backend-Core/database"
	"implude.kr/VOAH-Backend-Core/models"
	"implude.kr/VOAH-Backend-Core/utils/logger"
	"implude.kr/VOAH-Backend-Core/utils/validator"
)

type GetProfileRequest struct {
	UserID string `validate:"required,uuid4"`
}

func GetProfileCtrl(c *fiber.Ctx) error {

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

	if db.Model(&models.User{}).First(&user, uuid.MustParse(profileRequest.UserID)).Error != nil && !user.Visible {
		return c.Status(400).JSON(fiber.Map{
			"message": "User not found",
		})
	}
	log := logger.Logger
	log.Info(user.Displayname)
	userRoles := []models.Role{}
	if db.Model(&user).Association("Roles").Find(&userRoles) != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}
	userProjects := []models.Project{}
	if db.Model(&user).Association("Projects").Find(&userProjects) != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	user.Roles = userRoles
	fmt.Println(userRoles)
	fmt.Println(userProjects)
	user.Projects = userProjects

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
