package register

import (
	"context"
	"encoding/base64"
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"implude.kr/VOAH-Backend-Core/configs"
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
	if errArr := validator.ParseAndValidate(c, checkCodeRequest); errArr != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request body",
			"errArr":  errArr,
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
	var err error
	publicTeams := []models.Team{}
	if err = database.DB.Where(&models.Team{Public: true}).Find(&publicTeams).Error; err != nil && err != gorm.ErrRecordNotFound {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}
	myTeamInvites := []models.Invite{}
	if err = database.DB.Where(&models.Invite{RecieverEmail: checkCodeRequest.Email, TargetType: configs.TeamObject}).Find(&myTeamInvites).Error; err != nil && err != gorm.ErrRecordNotFound {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}
	invitedTeams := []models.Team{}
	for _, invite := range myTeamInvites {
		tempTeam := new(models.Team)
		if database.DB.First(tempTeam, invite.TargetID).Error != nil {
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

type SubmitCodeRequest struct {
	Code         string `json:"code" validate:"required,uuid"`
	Email        string `json:"email" validate:"required,email"`
	Password     string `json:"password" validate:"required,min=8,max=40,ascii"`
	Username     string `json:"username" validate:"required,min=1,max=30"`
	Displayname  string `json:"displayname" validate:"required,min=1,max=30"`
	Position     string `json:"position" validate:"max=30"`
	TeamID       string `json:"team-id" validate:"required,uuid4"`
	ProfileImage string `json:"profile-image" validate:"max=409600"`
}

func SubmitCodeCtrl(c *fiber.Ctx) error {
	submitCodeRequest := new(SubmitCodeRequest)
	if errArr := validator.ParseAndValidate(c, submitCodeRequest); errArr != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request body",
			"errArr":  errArr,
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
	if db.Where(&models.User{Username: submitCodeRequest.Username}).First(&checkUser).Error == nil {
		return c.Status(409).JSON(fiber.Map{
			"message": "Username already exists",
		})
	}

	// check if team is public
	foundTeam := new(models.Team)
	if db.First(&foundTeam, uuid.MustParse(submitCodeRequest.TeamID)).Error != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Team not found",
		})
	}
	if !foundTeam.Public {
		// check if user has invite
		checkInvite := new(models.Invite)
		if db.Where(&models.Invite{RecieverEmail: submitCodeRequest.Email, TargetID: uuid.MustParse(submitCodeRequest.TeamID), TargetType: configs.TeamObject}).First(&checkInvite).Error != nil {
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
		TeamID:      foundTeam.ID,
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

	userRole := models.Role{
		ID:    uuid.New(),
		Type:  "Personal",
		Name:  "[Personal]" + newUser.Username,
		Users: []models.User{*newUser},
	}

	if db.Create(&userRole).Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	userPermission := models.Permission{
		ID:     uuid.New(),
		Type:   configs.TeamObject,
		Target: foundTeam.ID,
		Scope:  configs.ReadPermissionScope,
		RoleID: userRole.ID,
	}

	if db.Create(&userPermission).Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Submit code success",
	})
}
