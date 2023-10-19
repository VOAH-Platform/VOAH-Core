package middleware

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"implude.kr/VOAH-Backend-Core/database"
)

var LastActivitMiddleware = func(c *fiber.Ctx) error {
	if c.Locals("user") == nil {
		return c.Next()
	}
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userID, err := uuid.Parse(claims["uuid"].(string))

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	// log current time to redis
	redis := database.Redis.LastActivityRedis
	ctx := context.Background()
	go redis.Set(ctx, userID.String(), time.Now().Unix(), time.Hour*24)

	return c.Next()
}
