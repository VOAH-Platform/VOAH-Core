package submitcode

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"implude.kr/VOAH-Backend-Core/configs"
	"implude.kr/VOAH-Backend-Core/database"
	"implude.kr/VOAH-Backend-Core/models"
	"implude.kr/VOAH-Backend-Core/utils/validator"
)

type SubmitCodeRequest struct {
	Code         string `json:"code" validate:"required,uuid"`
	Email        string `json:"email" validate:"required,email"`
	Password     string `json:"password" validate:"required,min=8,max=40,alphanumunicode"`
	Username     string `json:"username" validate:"required,min=1,max=30"`
	Displayname  string `json:"displayname" validate:"required,min=1,max=30"`
	Position     string `json:"position" validate:"max=30"`
	TeamID       string `json:"team-id" validate:"required,uuid4"`
	ProfileImage string `json:"profile-image" validate:"max=409600"`
}

func SubmitCodeCtrl(c *fiber.Ctx) error {
	submitCodeRequest := new(SubmitCodeRequest)
	if err := c.BodyParser(submitCodeRequest); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}
	if errArr := validator.VOAHValidator.Validate(submitCodeRequest); len(errArr) != 0 {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request",
			"error":   errArr,
		})
	}

	// check if unvalidated user exists
	ctx := context.Background()
	registerVerifyRedis := database.Redis.RegisterVerifyDB
	if registerVerifyRedis.Exists(ctx, submitCodeRequest.Code).Val() == 0 {
		return c.Status(401).JSON(fiber.Map{
			"message": "Invalid code",
		})
	} else if registerVerifyRedis.Get(ctx, submitCodeRequest.Code).Val() != submitCodeRequest.Email {
		return c.Status(401).JSON(fiber.Map{
			"message": "Invalid code",
		})
	}

	// check if username already exists
	db := database.DB
	checkUser := new(models.User)
	if err := db.Where(&models.User{Username: submitCodeRequest.Username}).First(&checkUser).Error; err == nil {
		return c.Status(409).JSON(fiber.Map{
			"message": "Username already exists",
		})
	}

	// check if team is public
	checkTeam := new(models.Team)
	if err := db.First(&checkTeam, submitCodeRequest.TeamID).Error; err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Team not found",
		})
	}
	if !checkTeam.Public {
		// check if user has invite
		checkInvite := new(models.Invite)
		if err := db.Where(&models.Invite{RecieverEmail: submitCodeRequest.Email, TargetID: uuid.MustParse(submitCodeRequest.TeamID), TargetType: models.TeamObject}).First(&checkInvite).Error; err != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": "Only invited user can join this team",
			})
		}
		db.Delete(&checkInvite)
	}

	// move unvalidated user to user
	pwHash, err := bcrypt.GenerateFromPassword([]byte(submitCodeRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}
	userID := uuid.New()

	newUser := &models.User{
		ID:          userID,
		Visible:     true,
		Email:       submitCodeRequest.Email,
		PWHash:      string(pwHash),
		Username:    submitCodeRequest.Username,
		Displayname: submitCodeRequest.Displayname,
		TeamID:      uuid.MustParse(submitCodeRequest.TeamID),
		Position:    submitCodeRequest.Position,
	}
	if registerVerifyRedis.Del(ctx, submitCodeRequest.Code).Err() != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}
	if db.Create(&newUser).Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	// save profile image
	if submitCodeRequest.ProfileImage != "" {
		decodedProfileImage, err := base64.StdEncoding.DecodeString(submitCodeRequest.ProfileImage)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{
				"message": "Bad Profile Image",
			})
		}
		// size check
		if len(decodedProfileImage) > 307200 {
			return c.Status(400).JSON(fiber.Map{
				"message": "Profile Image is too big",
			})
		}
		serverConf := configs.Env.Server
		// save profile image to ./data/profiles/{uuid}.png
		err = os.WriteFile(fmt.Sprintf(serverConf.DataDir+"/user-profiles/%s.png", userID.String()), decodedProfileImage, 0700)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"message": "Internal server error",
			})
		}
	}
	// set last activity and refresh
	lastActivityRedis := database.Redis.LastActivityRedis
	lastRefreshRedis := database.Redis.LastRefreshRedis
	go lastActivityRedis.Set(ctx, userID.String(), time.Now().Unix(), 0)
	go lastRefreshRedis.Set(ctx, userID.String(), time.Now().Unix(), 0)
	return c.JSON(fiber.Map{
		"message": "Submit code success",
	})
}
