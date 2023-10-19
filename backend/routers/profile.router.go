package routers

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"implude.kr/VOAH-Backend-Core/configs"
	"implude.kr/VOAH-Backend-Core/controllers/profile"
	"implude.kr/VOAH-Backend-Core/controllers/profile/profileimage"
	"implude.kr/VOAH-Backend-Core/middleware"
)

func addProfile(router *fiber.App) {
	profileGroup := router.Group("/api/profile") // profile router

	profileGroup.Use(
		jwtware.New(jwtware.Config{
			Filter: func(c *fiber.Ctx) bool {
				return c.Path() == "/api/profile/image" && c.Method() == "GET"
			},
			SigningKey: jwtware.SigningKey{Key: configs.Env.Auth.JWTSecret},
		}),
		middleware.LastActivitMiddleware,
	)
	profileGroup.Get("/", func(c *fiber.Ctx) error {
		return profile.GetProfileCtrl(c)
	})
	profileGroup.Post("/", func(c *fiber.Ctx) error {
		return profile.UpdateProfileCtrl(c)
	})
	profileGroup.Get("/image", func(c *fiber.Ctx) error {
		return profileimage.GetImageCtrl(c)
	})
	profileGroup.Post("/image", func(c *fiber.Ctx) error {
		return profileimage.UpdateImageCtrl(c)
	})
}
