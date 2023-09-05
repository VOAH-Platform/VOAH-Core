package routers

import (
	"github.com/gofiber/fiber/v2"
	"implude.kr/VOAH-Backend-Core/controllers/info"
)

func addInfo(router *fiber.App) {
	infoGroup := router.Group("/api/info") // info router

	infoGroup.Get("", func(c *fiber.Ctx) error {
		return info.GetInfoCtrl(c)
	})
}
