package routers

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"implude.kr/VOAH-Backend-Core/configs"
	"implude.kr/VOAH-Backend-Core/controllers/team/getteam"
	"implude.kr/VOAH-Backend-Core/controllers/team/teaminvite"
	"implude.kr/VOAH-Backend-Core/controllers/team/teamlist"
	"implude.kr/VOAH-Backend-Core/middleware"
)

func addTeam(router *fiber.App) {
	teamGroup := router.Group("/api/team") // team router

	teamGroup.Use(
		jwtware.New(jwtware.Config{
			SigningKey: jwtware.SigningKey{Key: configs.Env.Auth.JWTSecret},
		}),
		middleware.LastActivitMiddleware,
	)

	teamGroup.Get("/", func(c *fiber.Ctx) error {
		return getteam.GetTeamCtrl(c)
	})

	teamGroup.Get("/list", func(c *fiber.Ctx) error {
		return teamlist.TeamListCtrl(c)
	})
	teamGroup.Get("/invite", func(c *fiber.Ctx) error {
		return teaminvite.TeamInviteListCtrl(c)
	})
	teamGroup.Post("/invite/send", func(c *fiber.Ctx) error {
		return teaminvite.TeamInviteSendCtrl(c)
	})
}
