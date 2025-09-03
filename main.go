package main

import (
	"github.com/Sambit99/Basic-CRM-Go-SQLite/pkg/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	routes.RegisterLeadRoutes(app)

	app.Listen(":8080")
}
