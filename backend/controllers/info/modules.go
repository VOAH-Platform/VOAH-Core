package info

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"implude.kr/VOAH-Backend-Core/database"
	"implude.kr/VOAH-Backend-Core/models"
	"implude.kr/VOAH-Backend-Core/utils/validator"
)

func GetModuleList(c *fiber.Ctx) error {
	// get enabled modules
	db := database.DB

	var modules []models.Module

	if err := db.Where(&models.Module{Enabled: true}).Find(&modules).Error; err != nil && err != gorm.ErrRecordNotFound {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"modules": modules,
	})
}

type AddModuleRequest struct {
	ID               int `validate:"required"`
	Expose           bool
	Version          string `validate:"required" json:"version"`
	Name             string `validate:"required" json:"name"`
	Description      string `validate:"required" json:"description"`
	HostURL          string `validate:"required" json:"host-url"`
	HostInternalURL  string `validate:"required" json:"host-internal-url"`
	PermissionTypes  string `validate:"required" json:"permission-types"`
	PermissionScopes string `validate:"required" json:"permission-scopes"`
}

func AddModuleCtrl(c *fiber.Ctx) error {
	addModuleRequest := new(AddModuleRequest)
	if errArr := validator.ParseAndValidate(c, addModuleRequest); errArr != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request body",
			"errArr":  errArr,
		})
	}

	db := database.DB

	if err := db.Create(&models.Module{
		ID:               addModuleRequest.ID,
		Enabled:          true,
		Expose:           addModuleRequest.Expose,
		Version:          addModuleRequest.Version,
		Name:             addModuleRequest.Name,
		Description:      addModuleRequest.Description,
		HostURL:          addModuleRequest.HostURL,
		HostInternalURL:  addModuleRequest.HostInternalURL,
		PermissionTypes:  addModuleRequest.PermissionTypes,
		PermissionScopes: addModuleRequest.PermissionScopes,
	}).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
	})

}
