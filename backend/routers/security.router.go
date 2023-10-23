package routers

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"implude.kr/VOAH-Backend-Core/configs"
	"implude.kr/VOAH-Backend-Core/controllers/security"
	"implude.kr/VOAH-Backend-Core/middleware"
)

func addSecurity(router *fiber.App) {
	securityGroup := router.Group("/api/security") // profile router

	securityGroup.Use(
		jwtware.New(jwtware.Config{
			SigningKey: jwtware.SigningKey{Key: configs.Env.Auth.JWTSecret},
		}),
		middleware.LastActivitMiddleware,
	)

	securityGroup.Get("/2fa/add", func(c *fiber.Ctx) error {
		return security.Add2FACtrl(c)
	})
	securityGroup.Post("/2fa/validate", func(c *fiber.Ctx) error {
		return security.Validate2FACtrl(c)
	})
}
