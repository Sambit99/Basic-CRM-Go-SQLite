package routes

import (
	"github.com/Sambit99/Basic-CRM-Go-SQLite/pkg/controller"
	"github.com/gofiber/fiber/v2"
)

var RegisterLeadRoutes = func(app *fiber.App) {
	leadRouter := app.Group("/leads")

	leadRouter.Get("/", controller.GetLeads)
	leadRouter.Get("/:id", controller.GetLead)
	leadRouter.Post("/", controller.NewLead)
	leadRouter.Delete("/:id", controller.DeleteLead)
}
