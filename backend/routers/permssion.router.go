package routers

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"implude.kr/VOAH-Backend-Core/configs"
	"implude.kr/VOAH-Backend-Core/controllers/permission"
	"implude.kr/VOAH-Backend-Core/middleware"
)

func addPermission(router *fiber.App) {
	permissionGroup := router.Group("/api/permission") // profile router

	permissionGroup.Use(
		jwtware.New(jwtware.Config{
			SigningKey: jwtware.SigningKey{Key: configs.Env.Auth.JWTSecret},
		}),
		middleware.LastActivitMiddleware,
	)

	permissionGroup.Get("/", func(c *fiber.Ctx) error {
		return permission.GetMyPermissions(c)
	})
}
