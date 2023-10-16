package refresh

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"implude.kr/VOAH-Backend-Core/configs"
	"implude.kr/VOAH-Backend-Core/database"
	"implude.kr/VOAH-Backend-Core/utils/validator"
)

type RefreshRequest struct {
	RefreshToken string `json:"refresh-token" validate:"required,uuid4"`
	UserID       string `json:"user-id" validate:"required,uuid4"`
}

func RefreshCtrl(c *fiber.Ctx) error {
	refreshRequest := new(RefreshRequest)
	if errArr := validator.ParseAndValidate(c, refreshRequest); errArr != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request body",
			"errArr":  errArr,
		})
	}

	// check if refresh token exists
	redis := database.Redis.SessionRedis
	ctx := context.Background()

	if exist, _ := redis.SIsMember(ctx, refreshRequest.UserID, refreshRequest.RefreshToken).Result(); !exist {
		return c.Status(401).JSON(fiber.Map{
			"message": "Invalid refresh token",
		})
	}

	// generate new token
	authConf := configs.Env.Auth
	exp := time.Now().Add(time.Second * time.Duration(authConf.JWTExpire)).Unix()
	claims := jwt.MapClaims{
		"uuid": refreshRequest.UserID,
		"exp":  exp,
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(authConf.JWTSecret))

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}
	ctx = context.Background()
	lastRefreshRedis := database.Redis.LastRefreshRedis
	if err := lastRefreshRedis.Set(ctx, refreshRequest.UserID, time.Now().Unix(), 0).Err(); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	return c.JSON(fiber.Map{
		"message":      "Refresh success",
		"access-token": token,
		"exp":          exp,
	})

}
