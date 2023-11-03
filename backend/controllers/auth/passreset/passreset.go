package passreset

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"implude.kr/VOAH-Backend-Core/configs"
	"implude.kr/VOAH-Backend-Core/database"
	"implude.kr/VOAH-Backend-Core/utils/smtpsender"
	"implude.kr/VOAH-Backend-Core/utils/validator"
)

type passResetRequest struct {
	Email string `validate:"required,email"`
}

func PassResetCtrl(c *fiber.Ctx) error {
	passResetRequest := &passResetRequest{
		Email: c.Query("email"),
	}
	if errArr := validator.VOAHValidator.Validate(passResetRequest); len(errArr) != 0 {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request",
			"error":   errArr,
		})
	}

	// add redis
	ctx := context.Background()
	redis := database.Redis.PasswordResetRedis
	if redis.Exists(ctx, passResetRequest.Email).Val() == 1 {
		return c.Status(400).JSON(fiber.Map{
			"message": "Already requested",
		})
	}
	// send email

	smtpConf := configs.Env.SMTP
	authSetting := configs.Setting.Auth
	serverConf := configs.Env.Server
	code := uuid.New().String()

	mail := &smtpsender.Mail{
		From:    smtpConf.SystemAddress,
		Tos:     []string{passResetRequest.Email},
		Subject: authSetting.PasswordResetEmailSubject,
		Body:    strings.ReplaceAll(authSetting.PasswordResetEmailBody, "{{link}}", fmt.Sprintf("%s/auth/verify?type=passreset&email=%s&code=%s", serverConf.HostURL, passResetRequest.Email, code)),
	}

	go mail.ConnectAndSend()

	redis.Set(ctx, code, passResetRequest.Email, time.Duration(authSetting.PasswordResetExpire)*time.Second)
	return c.JSON(fiber.Map{
		"message": "Password reset email sent",
	})
}
