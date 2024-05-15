package main

import (
	"mysql-controller/pkg/routes"
	// "mysql-controller/pkg/env"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()

	// host := env.Get("UI_HOST", "localhost")
	// port := env.Get("UI_PORT", "3000")

	// app.Use(cors.New(cors.Config{
	// 	AllowOrigins:     "http://" + host + ":" + port,
	// 	AllowHeaders:     "Origin, Content-Type, Accept",
	// 	AllowCredentials: true,
	// }))

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowCredentials: false,
	}))

	routes.PublicRoutes(app)

	app.Listen(":3001")
}