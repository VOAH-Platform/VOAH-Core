package routers

import (
	"github.com/gofiber/fiber/v2"
	"implude.kr/VOAH-Backend-Core/controllers/auth/login"
	"implude.kr/VOAH-Backend-Core/controllers/auth/passreset"
	"implude.kr/VOAH-Backend-Core/controllers/auth/refresh"
	"implude.kr/VOAH-Backend-Core/controllers/auth/register"
)

func addAuth(router *fiber.App) {
	authGroup := router.Group("/api/auth") // auth router

	authGroup.Post("/login", func(c *fiber.Ctx) error {
		return login.LoginCtrl(c)
	})

	authGroup.Post("/register", func(c *fiber.Ctx) error {
		return register.RegisterCtrl(c)
	})
	authGroup.Post("/register/check", func(c *fiber.Ctx) error {
		return register.CheckCodeCtrl(c)
	})
	authGroup.Post("/register/submit", func(c *fiber.Ctx) error {
		return register.SubmitCodeCtrl(c)
	})
	authGroup.Post("/refresh", func(c *fiber.Ctx) error {
		return refresh.RefreshCtrl(c)
	})
	authGroup.Get("/passreset", func(c *fiber.Ctx) error {
		return passreset.PassResetCtrl(c)
	})
	authGroup.Post("/passreset/check", func(c *fiber.Ctx) error {
		return passreset.CheckPassResetCtrl(c)
	})
	authGroup.Post("/passreset/submit", func(c *fiber.Ctx) error {
		return passreset.SubmitPassResetCtrl(c)
	})
}
