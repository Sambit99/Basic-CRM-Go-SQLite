package controller

import (
	"strconv"

	"github.com/Sambit99/Basic-CRM-Go-SQLite/pkg/model"
	"github.com/gofiber/fiber/v2"
)

func GetLeads(c *fiber.Ctx) error {
	leads, err := model.GetLeads()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve leads",
		})
	}

	if len(leads) == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "No Leads Found",
		})
	}

	c.Set("Content-Type", "application/json")
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   leads,
	})
}

func GetLead(c *fiber.Ctx) error {
	id := c.Params("id")

	parsedId, err := strconv.ParseInt(id, 0, 0)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid ID Format",
		})
	}

	lead, err := model.GetLead(int(parsedId))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve lead",
		})
	}

	if lead.ID == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Lead Not found",
		})
	}

	c.Set("Content-Type", "application/json")

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data":   lead,
	})
}
