package main

import (
	"apisix-backend/config"
	"apisix-backend/route"
	"apisix-backend/share"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	config.LoadConfig()
	share.ConnectDB()
	share.MigrateDB()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: config.AppConfig.AllowOrigins,
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS,PATCH",
		AllowHeaders: "Content-Type, Authorization,apisix-key",
	}))
	route.SetupRoutes(app, config.AppConfig.APIPrefix)

	app.Listen(":8080")
}
