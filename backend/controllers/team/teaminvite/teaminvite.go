package teaminvite

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"implude.kr/VOAH-Backend-Core/database"
	"implude.kr/VOAH-Backend-Core/middleware"
	"implude.kr/VOAH-Backend-Core/models"
	"implude.kr/VOAH-Backend-Core/utils/permission"
	"implude.kr/VOAH-Backend-Core/utils/validator"
)

type TeamInviteRequest struct {
	TeamID    string `json:"team-id" validate:"required,uuid4"`
	Email     string `json:"email" validate:"required,email"`
	ExpireMin int    `json:"expire-min" validate:"required"`
}

func TeamInviteCtrl(c *fiber.Ctx) error {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}
	teamInviteRequest := new(TeamInviteRequest)
	if err := c.BodyParser(teamInviteRequest); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}
	if errArr := validator.VOAHValidator.Validate(teamInviteRequest); len(errArr) != 0 {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request",
			"error":   errArr,
		})
	}

	// get user
	foundUser := new(models.User)
	if err := database.DB.First(&foundUser, userID).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}
	// check if user has permission
	requierdPermission := []models.Permission{
		{
			Type:   models.TeamObject,
			Target: uuid.MustParse(teamInviteRequest.TeamID),
			Scope:  models.InvitePermissionScope,
		},
	}
	hasPerm, err := permission.UserPermissionCheck(*foundUser, requierdPermission)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	} else if !hasPerm {
		return c.Status(403).JSON(fiber.Map{
			"message": "Forbidden",
		})
	}
	//check team is private
	db := database.DB
	team := new(models.Team)
	if err := db.First(&team, uuid.MustParse(teamInviteRequest.TeamID)).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	} else if team.Public {
		return c.Status(403).JSON(fiber.Map{
			"message": "Invite is public team only",
		})
	}

	// create invite
	newInvite := models.Invite{
		ID:            uuid.New(),
		SenderID:      userID,
		RecieverEmail: teamInviteRequest.Email,
		TargetType:    models.TeamObject,
		TargetID:      uuid.MustParse(teamInviteRequest.TeamID),
		ExpireAt:      time.Now().Add(time.Minute * time.Duration(teamInviteRequest.ExpireMin)),
	}
	if err := db.Create(&newInvite).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Success",
	})
}
