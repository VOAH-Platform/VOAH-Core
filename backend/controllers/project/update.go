package project

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"implude.kr/VOAH-Backend-Core/configs"
	"implude.kr/VOAH-Backend-Core/database"
	"implude.kr/VOAH-Backend-Core/middleware"
	"implude.kr/VOAH-Backend-Core/models"
	"implude.kr/VOAH-Backend-Core/utils/checkperm"
	"implude.kr/VOAH-Backend-Core/utils/validator"
)

type UpdateProjectRequest struct {
	ProjectID   string `json:"project-id" validate:"required,uuid"`
	Public      bool   `json:"public"`
	Displayname string `json:"displayname" validate:"required,max=30"`
	Description string `json:"description" validate:"required,max=200"`
}

func UpdateProjectCtrl(c *fiber.Ctx) error {
	user, err := middleware.GetUserFromMiddleware(c)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	createRequest := new(UpdateProjectRequest)
	if errArr := validator.ParseAndValidate(c, createRequest); errArr != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request body",
			"errArr":  errArr,
		})
	}
	projectID := uuid.MustParse(createRequest.ProjectID)

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
			"message": "No Permission to Edit project",
		})
	}

	db := database.DB
	var project models.Project
	if err := db.Where(&models.Project{ID: projectID}).First(&project).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}
	project.Public = createRequest.Public
	project.Displayname = createRequest.Displayname
	project.Description = createRequest.Description
	if err := db.Save(&project).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Project updated",
	})
}
