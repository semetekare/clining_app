package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"clining_app/app/models"
	"clining_app/app/service"
)

type ClientController struct {
	clientService *service.ClientService
}

func NewClientController(clientService *service.ClientService) *ClientController {
	return &ClientController{clientService: clientService}
}

func (c *ClientController) CreateClient(ctx *fiber.Ctx) error {
	var client models.Client
	if err := ctx.BodyParser(&client); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}
	if err := c.clientService.CreateClient(&client); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(fiber.StatusCreated).JSON(client)
}

func (c *ClientController) GetClient(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}
	client, err := c.clientService.GetClientByID(uint(id))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Client not found"})
	}
	return ctx.JSON(client)
}

func (c *ClientController) GetAllClients(ctx *fiber.Ctx) error {
	clients, err := c.clientService.GetAllClients()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(clients)
}

func (c *ClientController) UpdateClient(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}
	var client models.Client
	if err := ctx.BodyParser(&client); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}
	client.ID = uint(id)
	if err := c.clientService.UpdateClient(&client); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(client)
}

func (c *ClientController) DeleteClient(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}
	if err := c.clientService.DeleteClient(uint(id)); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.SendStatus(fiber.StatusNoContent)
}