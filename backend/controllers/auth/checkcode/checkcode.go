package checkcode

import (
	"context"
	"errors"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"implude.kr/VOAH-Backend-Core/database"
	"implude.kr/VOAH-Backend-Core/models"
	"implude.kr/VOAH-Backend-Core/utils/validator"
)

type CheckCodeRequest struct {
	Email string `json:"email" validate:"required,email"`
	Code  string `json:"code" validate:"required,uuid"`
}

func CheckCodeCtrl(c *fiber.Ctx) error {
	checkCodeRequest := new(CheckCodeRequest)
	if err := c.BodyParser(checkCodeRequest); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}

	if errArr := validator.VOAHValidator.Validate(checkCodeRequest); len(errArr) != 0 {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request",
			"error":   errArr,
		})
	}

	// check if unvalidated user exists
	ctx := context.Background()
	redis := database.Redis.RegisterVerifyDB
	if redis.Exists(ctx, checkCodeRequest.Code).Val() == 0 {
		return c.Status(401).JSON(fiber.Map{
			"message": "Invalid code",
		})
	} else if redis.Get(ctx, checkCodeRequest.Code).Val() != checkCodeRequest.Email {
		return c.Status(401).JSON(fiber.Map{
			"message": "Invalid code",
		})
	}

	publicTeams := []models.Team{}
	if err := database.DB.Where(&models.Team{Public: true}).Find(&publicTeams).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}
	myTeamInvites := []models.Invite{}
	if err := database.DB.Where(&models.Invite{RecieverEmail: checkCodeRequest.Email, TargetType: models.TeamObject}).Find(&myTeamInvites).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}
	invitedTeams := []models.Team{}
	for _, invite := range myTeamInvites {
		tempTeam := new(models.Team)
		if err := database.DB.First(tempTeam, invite.TargetID).Error; err != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": "Internal server error",
			})
		}
		invitedTeams = append(invitedTeams, *tempTeam)
	}

	return c.JSON(fiber.Map{
		"message":       "Valid code",
		"public-teams":  publicTeams,
		"invited-teams": invitedTeams,
	})
}
