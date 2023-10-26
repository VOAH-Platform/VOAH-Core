package permission

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"implude.kr/VOAH-Backend-Core/database"
	"implude.kr/VOAH-Backend-Core/models"
	"implude.kr/VOAH-Backend-Core/utils/checkperm"
	"implude.kr/VOAH-Backend-Core/utils/validator"
)

func GetPersonalPermission(c *fiber.Ctx) error {
	userID, err := uuid.Parse(c.Query("user-id", ""))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid user id",
		})
	}
	// get user with id
	var user models.User
	db := database.DB
	if err = db.Where(&models.User{ID: userID}).First(&user).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "User not found",
		})
	}
	// get user roles
	var roles []models.Role
	roles, err = checkperm.GetUserRoleArr(&user)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}
	// get user personal role
	var findFlag bool = false
	var personalRole models.Role
	for _, role := range roles {
		if role.Type == "Personal" {
			findFlag = true
			personalRole = role
			break
		}
	}
	if !findFlag {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}
	// get user permissions
	var permissions []models.Permission
	permissions, err = checkperm.GetPermissionByRoleArr([]models.Role{personalRole})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"perms":   permissions,
	})
}

type AddPersonalPermRequest struct {
	UserID     uuid.UUID `json:"user_id"`
	TargetType string    `json:"target_type"`
	TargetID   uuid.UUID `json:"target_id"`
	Scope      string    `json:"scope"`
}

func AddPersonalPermCtrl(c *fiber.Ctx) error {
	var addPermRequset AddPersonalPermRequest
	if errArr := validator.ParseAndValidate(c, &addPermRequset); errArr != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request body",
			"details": errArr,
		})
	}

	// user, err := middleware.GetUserFromMiddleware(c)
	// if err != nil {
	// 	return c.Status(500).JSON(fiber.Map{
	// 		"message": "Internal server error",
	// 	})
	// }
	return nil
}
