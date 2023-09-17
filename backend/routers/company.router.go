package routers

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"implude.kr/VOAH-Backend-Core/configs"
	"implude.kr/VOAH-Backend-Core/controllers/company"
	"implude.kr/VOAH-Backend-Core/middleware"
)

func addCompany(router *fiber.App) {
	companyGroup := router.Group("/api/company") // company router

	companyGroup.Use(
		jwtware.New(jwtware.Config{
			SigningKey: jwtware.SigningKey{Key: configs.Env.Auth.JWTSecret},
		}),
		middleware.LastActivitMiddleware,
	)

	companyGroup.Get("/", func(c *fiber.Ctx) error {
		return company.GetCompanyInfo(c)
	})
	companyGroup.Get("/image", func(c *fiber.Ctx) error {
		return company.GetCompanyImage(c)
	})
}
