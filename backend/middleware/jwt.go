package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func GetUserIDFromMiddleware(c *fiber.Ctx) (userID uuid.UUID, err error) {
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userID, err = uuid.Parse(claims["uuid"].(string))

	return
}
