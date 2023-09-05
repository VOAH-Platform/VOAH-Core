package register

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"implude.kr/VOAH-Backend-Core/configs"
	"implude.kr/VOAH-Backend-Core/database"
	"implude.kr/VOAH-Backend-Core/models"
	"implude.kr/VOAH-Backend-Core/utils/smtpsender"
	"implude.kr/VOAH-Backend-Core/utils/validator"
)

type RegisterRequest struct {
	Email string `json:"email" validate:"required,email"`
}

func RegisterCtrl(c *fiber.Ctx) error {

	authSetting := configs.Setting.Auth

	// if AllowRegister is turned off in config, return 403
	if !authSetting.AllowRegister {
		return c.Status(403).JSON(fiber.Map{
			"message": "Register is not allowed",
		})
	}
	registerRequest := new(RegisterRequest)
	if err := c.BodyParser(registerRequest); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request",
		})
	}
	if errArr := validator.VOAHValidator.Validate(registerRequest); len(errArr) != 0 {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request",
			"error":   errArr,
		})
	}

	// if AllowOnlyDomain is turned on in config, check if email is from allowed domain
	if authSetting.AllowOnlyDomain {
		if strings.Split(registerRequest.Email, "@")[1] != configs.Setting.Company.Domain {
			return c.Status(400).JSON(fiber.Map{
				"message": "Invalid request",
				"error":   "Email is not from allowed domain",
			})
		}
	}

	// check if email or username already exists
	checkUser := new(models.User)
	db := database.DB
	db.Where(&models.User{Email: registerRequest.Email}).First(&checkUser)
	if checkUser.ID != uuid.Nil {
		return c.Status(409).JSON(fiber.Map{
			"message": "Email or Username already exists",
		})
	}

	// set unvalidated user on redis
	ctx := context.Background()
	redis := database.Redis.RegisterVerifyDB

	verifcationCode := uuid.New().String()
	redis.Set(ctx, verifcationCode, registerRequest.Email, time.Minute*time.Duration(authSetting.EmailVerificattionExpire))

	smtpConf := configs.Env.SMTP
	serverConf := configs.Env.Server

	mail := &smtpsender.Mail{
		From:    smtpConf.SystemAddress,
		Tos:     []string{registerRequest.Email},
		Subject: authSetting.VerificationEmailSubject,
		Body:    strings.ReplaceAll(authSetting.VerificationEmailBody, "{{link}}", fmt.Sprintf("%s/auth/verify?type=email&user=%s&code=%s", serverConf.HostURL, registerRequest.Email, verifcationCode)),
	}

	go mail.ConnectAndSend()

	return c.JSON(fiber.Map{
		"message": "Send verification code to email",
	})
}
