package routers

import (
	"github.com/gofiber/fiber/v2"
	"implude.kr/VOAH-Backend-Core/controllers/auth/checkcode"
	"implude.kr/VOAH-Backend-Core/controllers/auth/login"
	"implude.kr/VOAH-Backend-Core/controllers/auth/refresh"
	"implude.kr/VOAH-Backend-Core/controllers/auth/register"
	"implude.kr/VOAH-Backend-Core/controllers/auth/submitcode"
)

func addAuth(router *fiber.App) {
	authGroup := router.Group("/api/auth") // auth router

	authGroup.Post("/login", func(c *fiber.Ctx) error {
		return login.LoginCtrl(c)
	})

	authGroup.Post("/register", func(c *fiber.Ctx) error {
		return register.RegisterCtrl(c)
	})

	authGroup.Post("/refresh", func(c *fiber.Ctx) error {
		return refresh.RefreshCtrl(c)
	})
	authGroup.Post("/checkcode", func(c *fiber.Ctx) error {
		return checkcode.CheckCodeCtrl(c)
	})
	authGroup.Post("/submitcode", func(c *fiber.Ctx) error {
		return submitcode.SubmitCodeCtrl(c)
	})
}
