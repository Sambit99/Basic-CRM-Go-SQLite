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

func NewLead(c *fiber.Ctx) error {
	var lead model.Lead

	if err := c.BodyParser(&lead); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "failed",
			"message": "Error while parsing body",
			"error":   err.Error(),
		})
	}

	newLead, err := model.NewLead(lead)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "failed",
			"message": "Error while creating New Lead",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status": "success",
		"data":   newLead,
	})
}

func DeleteLead(c *fiber.Ctx) error {
	id := c.Params("id")

	leadId, err := strconv.ParseInt(id, 0, 0)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "failed",
			"message": "Error while parsing ID",
		})
	}

	_, err = model.DeleteLead(int(leadId))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "failed",
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
		"status":  "success",
		"message": "Lead Deleted Successfully",
	})
}
