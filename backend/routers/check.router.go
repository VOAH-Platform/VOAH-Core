package routers

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"implude.kr/VOAH-Backend-Core/configs"
	"implude.kr/VOAH-Backend-Core/controllers/check"
	"implude.kr/VOAH-Backend-Core/middleware"
)

func addCheck(router *fiber.App) {
	checkGroup := router.Group("/api/check") // profile router

	checkGroup.Use(
		middleware.CheckAPIKeyMiddleware,
		jwtware.New(jwtware.Config{
			SigningKey: jwtware.SigningKey{Key: configs.Env.Auth.JWTSecret},
		}),
		middleware.LastActivitMiddleware,
	)
	checkGroup.Get("/", func(c *fiber.Ctx) error {
		return check.CheckUserCtrl(c)
	})
}
