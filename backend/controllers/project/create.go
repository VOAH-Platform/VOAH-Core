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

type CreateProjectRequest struct {
	Public      bool   `json:"public"`
	Displayname string `json:"displayname" validate:"required,max=30"`
	Description string `json:"description" validate:"required,max=200"`
}

func CreateProjectCtrl(c *fiber.Ctx) error {
	user, err := middleware.GetUserFromMiddleware(c)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}
	hasPerm, err := checkperm.UserPermissionCheck(user, []models.Permission{
		{
			Type:   configs.CompanyObject,
			Scope:  configs.EditPermissionScope,
			Target: configs.CompanyID,
		},
	})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}
	if !hasPerm {
		return c.Status(403).JSON(fiber.Map{
			"message": "No Permission to create project",
		})
	}
	createRequest := new(CreateProjectRequest)
	if errArr := validator.ParseAndValidate(c, createRequest); errArr != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request body",
			"errArr":  errArr,
		})
	}
	project := &models.Project{
		ID:          uuid.New(),
		Public:      createRequest.Public,
		Displayname: createRequest.Displayname,
		Description: createRequest.Description,
		Users:       []models.User{*user},
	}
	// get user personal role
	db := database.DB
	userRoles, err := checkperm.GetUserRoleArr(user)
	if err != nil {
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

	permission := &models.Permission{
		ID:     uuid.New(),
		Type:   configs.ProjectObject,
		Target: project.ID,
		Scope:  configs.AdminPermissionScope,
		RoleID: userRole.ID,
	}

	if err := db.Create(project).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}
	if err := db.Create(permission).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}
	return c.Status(200).JSON(fiber.Map{
		"success":    true,
		"project-id": project.ID.String(),
	})
}
