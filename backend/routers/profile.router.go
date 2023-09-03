package routers

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"implude.kr/VOAH-Backend-Core/configs"
	"implude.kr/VOAH-Backend-Core/controllers/profile/getimage"
	"implude.kr/VOAH-Backend-Core/controllers/profile/getprofile"
	"implude.kr/VOAH-Backend-Core/controllers/profile/updateimage"
	"implude.kr/VOAH-Backend-Core/controllers/profile/updateprofile"
	"implude.kr/VOAH-Backend-Core/middleware"
)

func addProfile(router *fiber.App) {
	profileGroup := router.Group("/api/profile") // profile router

	profileGroup.Use(
		jwtware.New(jwtware.Config{
			SigningKey: jwtware.SigningKey{Key: configs.Env.Auth.JWTSecret},
		}),
		middleware.LastActivitMiddleware,
	)
	profileGroup.Get("/", func(c *fiber.Ctx) error {
		return getprofile.GetProfileCtrl(c)
	})
	profileGroup.Post("/", func(c *fiber.Ctx) error {
		return updateprofile.UpdateProfileCtrl(c)
	})
	profileGroup.Get("/image", func(c *fiber.Ctx) error {
		return getimage.GetImageCtrl(c)
	})
	profileGroup.Post("/image", func(c *fiber.Ctx) error {
		return updateimage.UpdateImageCtrl(c)
	})
}
