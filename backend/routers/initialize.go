package routers

import "github.com/gofiber/fiber/v2"

func Initialize(router *fiber.App) {
	addAuth(router)
	addInfo(router)
	addProfile(router)
	addTeam(router)
	addServer(router)
	addCompany(router)
	addRole(router)
	addPermission(router)
	addSecurity(router)
	addProject(router)
}
