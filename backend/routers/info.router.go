package routers

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"implude.kr/VOAH-Backend-Core/configs"
	"implude.kr/VOAH-Backend-Core/controllers/info"
	"implude.kr/VOAH-Backend-Core/middleware"
)

func addInfo(router *fiber.App) {
	infoGroup := router.Group("/api/info") // info router

	infoGroup.Get("", func(c *fiber.Ctx) error {
		return info.GetInfoCtrl(c)
	})
	infoModuleGroup := infoGroup.Group("/modules")
	infoModuleGroup.Use(
		jwtware.New(jwtware.Config{
			SigningKey: jwtware.SigningKey{Key: configs.Env.Auth.JWTSecret},
		}),
		middleware.LastActivitMiddleware,
	)
	infoModuleGroup.Get("", func(c *fiber.Ctx) error {
		return info.GetModuleList(c)
	})
	infoModuleGroup.Post("", func(c *fiber.Ctx) error {
		return info.AddModuleCtrl(c)
	})
}
