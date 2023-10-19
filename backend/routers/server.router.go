package routers

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"implude.kr/VOAH-Backend-Core/configs"
	"implude.kr/VOAH-Backend-Core/controllers/server"
	"implude.kr/VOAH-Backend-Core/middleware"
)

func addServer(router *fiber.App) {
	serverGroup := router.Group("/api/server") // profile router

	serverGroup.Use(
		middleware.CheckAPIKeyMiddleware,
		jwtware.New(jwtware.Config{
			SigningKey: jwtware.SigningKey{Key: configs.Env.Auth.JWTSecret},
		}),
		middleware.LastActivitMiddleware,
	)
	serverGroup.Get("/user", func(c *fiber.Ctx) error {
		return server.CheckUserCtrl(c)
	})
}
