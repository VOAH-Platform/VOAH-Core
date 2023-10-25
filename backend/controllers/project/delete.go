package project

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"implude.kr/VOAH-Backend-Core/configs"
	"implude.kr/VOAH-Backend-Core/database"
	"implude.kr/VOAH-Backend-Core/middleware"
	"implude.kr/VOAH-Backend-Core/models"
	"implude.kr/VOAH-Backend-Core/utils/checkperm"
)

func DeleteProjectCtrl(c *fiber.Ctx) error {
	user, err := middleware.GetUserFromMiddleware(c)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}
	projectID, err := uuid.Parse(c.Query("project-id", ""))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid projectID",
		})
	}

	hasPerm, err := checkperm.UserPermissionCheck(user, []models.Permission{
		{
			Type:   configs.ProjectObject,
			Scope:  configs.AdminPermissionScope,
			Target: projectID,
		},
	})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	} else if !hasPerm {
		return c.Status(403).JSON(fiber.Map{
			"message": "No Permission to delete project",
		})
	}

	db := database.DB
	var project models.Project
	var permissions []models.Permission
	if db.Where(&models.Project{ID: projectID}).First(&project).Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}
	if db.Delete(&project).Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}
	if db.Where(&models.Permission{Type: configs.ProjectObject, Target: projectID}).Find(&permissions).Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}
	if db.Delete(&permissions).Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}
	return c.JSON(fiber.Map{
		"message": "success",
	})
}
