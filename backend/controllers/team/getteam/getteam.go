package getteam

import (
	"github.com/gofiber/fiber/v2"
	"implude.kr/VOAH-Backend-Core/database"
	"implude.kr/VOAH-Backend-Core/models"
	"implude.kr/VOAH-Backend-Core/utils/validator"
)

type GetTeamRequest struct {
	TeamID string `json:"team-id" validate:"required,uuid4"`
}

func GetTeamCtrl(c *fiber.Ctx) error {
	getTeamRequest := &GetTeamRequest{
		TeamID: c.Params("team-id"),
	}
	if errArr := validator.VOAHValidator.Validate(getTeamRequest); len(errArr) != 0 {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request",
			"error":   errArr,
		})
	}

	// get team
	db := database.DB
	foundTeam := new(models.Team)
	if err := db.First(&foundTeam, getTeamRequest.TeamID).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Team not found",
		})
	}
	return c.JSON(fiber.Map{
		"message": "Success",
		"team":    foundTeam,
	})
}
