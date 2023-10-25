package routers

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"implude.kr/VOAH-Backend-Core/configs"
	"implude.kr/VOAH-Backend-Core/controllers/project"
	"implude.kr/VOAH-Backend-Core/middleware"
)

func addProject(router *fiber.App) {
	projectGroup := router.Group("/api/project")
	projectGroup.Use(
		jwtware.New(jwtware.Config{
			SigningKey: jwtware.SigningKey{Key: configs.Env.Auth.JWTSecret},
		}),
		middleware.LastActivitMiddleware,
	)
	projectGroup.Post("/", func(c *fiber.Ctx) error {
		return project.CreateProjectCtrl(c)
	})
	projectGroup.Post("/update", func(c *fiber.Ctx) error {
		return project.UpdateProjectCtrl(c)
	})
	projectGroup.Delete("/", func(c *fiber.Ctx) error {
		return project.DeleteProjectCtrl(c)
	})
}
