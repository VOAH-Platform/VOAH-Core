package routers

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"implude.kr/VOAH-Backend-Core/configs"
	"implude.kr/VOAH-Backend-Core/controllers/role"
	"implude.kr/VOAH-Backend-Core/middleware"
)

func addRole(router *fiber.App) {
	roleGroup := router.Group("/api/role") // profile router

	roleGroup.Use(
		jwtware.New(jwtware.Config{
			SigningKey: jwtware.SigningKey{Key: configs.Env.Auth.JWTSecret},
		}),
		middleware.LastActivitMiddleware,
	)

	roleGroup.Get("/", func(c *fiber.Ctx) error {
		return role.GetMyRoles(c)
	})
}
