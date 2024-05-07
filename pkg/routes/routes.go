package routes

import (
	"github.com/gofiber/fiber/v2"

	"mysql-controller/pkg/commands"
)

func PublicRoutes(a *fiber.App) {
	route := a.Group("/api/v1")

	route.Get("/getall", commands.GetAll)
    route.Get("/get/:id", commands.Get)
    route.Post("/insert/:taskName", commands.Insert)
    route.Post("/delete/:id", commands.Delete)
    route.Post("/update/:id.:status", commands.Update)
}
