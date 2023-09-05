package main

import (
	"fmt"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"implude.kr/VOAH-Backend-Core/configs"
	"implude.kr/VOAH-Backend-Core/database"
	"implude.kr/VOAH-Backend-Core/routers"
	"implude.kr/VOAH-Backend-Core/utils/directory"
	"implude.kr/VOAH-Backend-Core/utils/logger"
	"implude.kr/VOAH-Backend-Core/utils/module"
	"implude.kr/VOAH-Backend-Core/utils/rootuser"
	"implude.kr/VOAH-Backend-Core/utils/smtpsender"
)

func main() {
	configs.LoadEnv()     // Load configs
	configs.LoadSetting() // Load settings
	logger.InitLogger()   // Intitialize logger

	var wait sync.WaitGroup
	wait.Add(5)
	go database.ConnectDB(&wait) // Connect to database
	go database.InitRedis(&wait)
	go directory.InitDirectory(&wait)
	go smtpsender.InitSMTP(&wait)
	go configs.LoadAPIKey(&wait)
	wait.Wait()

	rootuser.InitRootUser()
	module.InitModules()

	serverConf := configs.Env.Server
	log := logger.Logger

	app := fiber.New()

	// CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: serverConf.CSRFOrigin,
		AllowHeaders: "*",
		AllowMethods: "*",
	}))

	routers.Initialize(app) // Initialize router

	// Static Files
	app.Get("/logo.svg", func(c *fiber.Ctx) error {
		return c.SendFile("./public/logo.svg")
	})
	app.Static("/assets", "./public/assets")
	app.Static("*", "./public/index.html")

	log.Fatal(app.Listen(fmt.Sprintf("%s:%d", serverConf.Host, serverConf.Port)))
}
