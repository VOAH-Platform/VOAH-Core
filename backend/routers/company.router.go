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
			Filter: func(c *fiber.Ctx) bool {
				return c.Path() == "/api/company/image" && c.Method() == "GET"
			},
			SigningKey: jwtware.SigningKey{Key: configs.Env.Auth.JWTSecret},
		}),
		middleware.LastActivitMiddleware,
	)

	companyGroup.Get("/", func(c *fiber.Ctx) error {
		return company.GetCompanyInfo(c)
	})

	companySetting := configs.Setting.Company
	companyGroup.Static("/image", configs.Env.Server.DataDir+"/"+companySetting.LogoImageRelativePath)
}
