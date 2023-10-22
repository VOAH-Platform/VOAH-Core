package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"implude.kr/VOAH-Backend-Core/database"
	"implude.kr/VOAH-Backend-Core/models"
)

func GetUserFromMiddleware(c *fiber.Ctx) (*models.User, error) {
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)

	userID, err := uuid.Parse(claims["uuid"].(string))
	if err != nil {
		return nil, err
	}
	db := database.DB
	foundUser := &models.User{}
	if db.Where(&models.User{ID: userID}).First(&foundUser).Error != nil {
		return nil, err
	}
	return foundUser, nil
}

func GetUserIDFromMiddleware(c *fiber.Ctx) (userID uuid.UUID, err error) {
	token := c.Locals("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userID, err = uuid.Parse(claims["uuid"].(string))
	return
}
