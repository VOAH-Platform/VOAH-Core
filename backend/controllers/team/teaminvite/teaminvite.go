package teaminvite

import (
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"implude.kr/VOAH-Backend-Core/configs"
	"implude.kr/VOAH-Backend-Core/database"
	"implude.kr/VOAH-Backend-Core/middleware"
	"implude.kr/VOAH-Backend-Core/models"
	"implude.kr/VOAH-Backend-Core/utils/async"
	"implude.kr/VOAH-Backend-Core/utils/permission"
	"implude.kr/VOAH-Backend-Core/utils/validator"
)

func TeamInviteListCtrl(c *fiber.Ctx) error {
	user, err := middleware.GetUserFromMiddleware(c)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}
	// get all team invites on user
	db := database.DB
	receivedInvites := new([]models.Invite)
	sentInvites := new([]models.Invite)

	var wait sync.WaitGroup
	wait.Add(2)
	async.AsyncDBQuery(func() *gorm.DB {
		return db.Where(&models.Invite{RecieverEmail: user.Email, TargetType: configs.TeamObject}).Find(&receivedInvites)
	}, &wait)
	async.AsyncDBQuery(func() *gorm.DB {
		return db.Where(&models.Invite{SenderID: user.ID, TargetType: configs.TeamObject}).Find(&sentInvites)
	}, &wait)
	wait.Wait()

	return c.JSON(fiber.Map{
		"message":          "Success",
		"sent-invites":     sentInvites,
		"received-invites": receivedInvites,
	})
}

type TeamInviteSendRequest struct {
	TeamID    string `json:"team-id" validate:"required,uuid4"`
	Email     string `json:"email" validate:"required,email"`
	ExpireMin int    `json:"expire-min" validate:"required"`
}

func TeamInviteSendCtrl(c *fiber.Ctx) error {
	user, err := middleware.GetUserFromMiddleware(c)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}
	teamInviteRequest := new(TeamInviteSendRequest)
	if errArr := validator.ParseAndValidate(c, teamInviteRequest); errArr != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request body",
			"errArr":  errArr,
		})
	}
	// check if user has permission
	requierdPermission := []models.Permission{
		{
			Type:   configs.TeamObject,
			Target: uuid.MustParse(teamInviteRequest.TeamID),
			Scope:  configs.InvitePermissionScope,
		},
	}
	hasPerm, err := permission.UserPermissionCheck(user, requierdPermission)
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
		SenderID:      user.ID,
		RecieverEmail: teamInviteRequest.Email,
		TargetType:    configs.TeamObject,
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

type TeamInviteAcceptRequest struct {
	InviteID string `json:"invite-id" validate:"required,uuid4"`
}
