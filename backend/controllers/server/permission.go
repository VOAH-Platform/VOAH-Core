package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"implude.kr/VOAH-Backend-Core/configs"
	"implude.kr/VOAH-Backend-Core/database"
	"implude.kr/VOAH-Backend-Core/models"
	"implude.kr/VOAH-Backend-Core/utils/validator"
)

type InjectPermissionRequest struct {
	Type     string `json:"type"`
	Scope    string `json:"scope"`
	TargetID string `json:"target-id"`
	UserID   string `json:"user-id"`
}

func InjectPermissionToUserCtrl(c *fiber.Ctx) error {
	injectRequest := new(InjectPermissionRequest)
	if errArr := validator.ParseAndValidate(c, injectRequest); errArr != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request body",
			"errArr":  errArr,
		})
	}
	// find user
	db := database.DB
	user := &models.User{}
	if err := db.First(&user, uuid.MustParse(injectRequest.UserID)).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "User not found",
		})
	}
	// find user personal role
	userRoles := []models.Role{}
	if err := db.Model(&user).Association("Roles").Find(&userRoles); err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}
	var findFlag bool = false
	var userRole models.Role
	for _, role := range userRoles {
		if role.Type == "Personal" {
			findFlag = true
			userRole = role
		}
	}
	if !findFlag {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}
	// add permission to user personal role
	permission := &models.Permission{
		ID:     uuid.New(),
		Type:   configs.ObjectType(injectRequest.Type),
		Target: uuid.MustParse(injectRequest.TargetID),
		Scope:  configs.PermissionScope(injectRequest.Scope),
		RoleID: userRole.ID,
	}
	if err := db.Create(permission).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Success",
	})

}
