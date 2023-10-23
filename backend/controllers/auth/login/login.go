package login

import (
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/pquerna/otp/totp"
	"golang.org/x/crypto/bcrypt"
	"implude.kr/VOAH-Backend-Core/configs"
	"implude.kr/VOAH-Backend-Core/database"
	"implude.kr/VOAH-Backend-Core/models"
	"implude.kr/VOAH-Backend-Core/utils/validator"
)

type LoginRequest struct {
	Email        string            `json:"email" validate:"required,email"`
	Password     string            `json:"password" validate:"required,min=8,max=40"`
	DeviceID     string            `json:"device-id" validate:"required,uuid4"`
	DeviceType   models.DeviceType `json:"device-type" validate:"required,min=1,max=6"`
	DeviceDetail string            `json:"device-detail" validate:"required,min=1,max=30"`
	TwoFACode    string            `json:"2fa-code"`
}

func LoginCtrl(c *fiber.Ctx) error {
	loginRequest := new(LoginRequest)
	if errArr := validator.ParseAndValidate(c, loginRequest); errArr != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid request body",
			"errArr":  errArr,
		})
	}

	db := database.DB
	user := new(models.User)
	db.Where(&models.User{Email: loginRequest.Email}).First(&user)
	if bcrypt.CompareHashAndPassword([]byte(user.PWHash), []byte(loginRequest.Password)) != nil {
		return c.Status(401).JSON(fiber.Map{
			"message": "Invalid Email or Password",
		})
	} else if user.TwoFA && !totp.Validate(loginRequest.TwoFACode, user.TwoFAKey) {
		return c.Status(498).JSON(fiber.Map{
			"message": "Invalid 2FA Code",
		})
	}
	// check sessionID is conflict
	foundSession := new(models.Session)
	if err := db.Where(&models.Session{DeviceID: uuid.MustParse(loginRequest.DeviceID)}).First(&foundSession).Error; err == nil {
		return c.Status(409).JSON(fiber.Map{
			"message": "Device is conflict",
		})
	}

	// create session
	refreshToken := uuid.New()
	session := &models.Session{
		ID:           refreshToken,
		UserID:       user.ID,
		DeviceID:     uuid.MustParse(loginRequest.DeviceID),
		DeviceType:   loginRequest.DeviceType,
		DeviceDetail: loginRequest.DeviceDetail,
	}
	db.Create(&session)

	// enroll refresh token to redis
	redis := database.Redis.SessionRedis
	ctx := context.Background()
	redis.SAdd(ctx, user.ID.String(), refreshToken.String())

	authConf := configs.Env.Auth

	exp := time.Now().Add(time.Second * time.Duration(authConf.JWTExpire)).Unix()
	claims := jwt.MapClaims{
		"uuid": user.ID,
		"exp":  exp,
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(authConf.JWTSecret))

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"message": "Internal server error",
		})
	}

	return c.JSON(fiber.Map{
		"message":       "Login success",
		"user-id":       user.ID,
		"refresh-token": refreshToken,
		"access-token":  token,
		"exp":           exp,
	})
}
