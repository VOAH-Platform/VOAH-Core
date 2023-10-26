package routers

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"implude.kr/VOAH-Backend-Core/configs"
	"implude.kr/VOAH-Backend-Core/controllers/permission"
	"implude.kr/VOAH-Backend-Core/middleware"
	"implude.kr/VOAH-Backend-Core/wshandler"
)

func addPermission(router *fiber.App) {
	permissionGroup := router.Group("/api/permission") // profile router

	permissionGroup.Use(
		jwtware.New(jwtware.Config{
			Filter: func(c *fiber.Ctx) bool {
				return c.Method() == "GET"
			},
			SigningKey: jwtware.SigningKey{Key: configs.Env.Auth.JWTSecret},
		}),
		middleware.LastActivitMiddleware,
	)

	permissionGroup.Get("/", func(c *fiber.Ctx) error {
		return permission.GetMyPermissions(c)
	})
	permissionGroup.Get("/personal", func(c *fiber.Ctx) error {
		return permission.GetPersonalPermission(c)
	})
	permWSGroup := router.Group("/api/ws/permission")
	permWSGroup.Get("/", websocket.New(wshandler.PermissionWebsocket()))
}
