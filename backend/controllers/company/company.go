package company

import (
	"github.com/gofiber/fiber/v2"
	"implude.kr/VOAH-Backend-Core/configs"
)

func GetCompanyInfo(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"success": true,
		"company": configs.Setting.Company,
	})
}

func GetCompanyImage(c *fiber.Ctx) error {
	// _, err := middleware.GetUserFromMiddleware(c)
	// if err != nil {
	// 	return c.Status(401).JSON(fiber.Map{
	// 		"success": false,
	// 	})
	// }
	serverConf := configs.Env.Server
	companySetting := configs.Setting.Company
	return c.SendFile(serverConf.DataDir + "/" + companySetting.LogoImageRelativePath)
}
