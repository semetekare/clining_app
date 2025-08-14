package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"clining_app/app/models"
	"clining_app/app/service"
)

type ServicePriceController struct {
	servicePriceService *service.ServicePriceService
}

func NewServicePriceController(servicePriceService *service.ServicePriceService) *ServicePriceController {
	return &ServicePriceController{servicePriceService: servicePriceService}
}

func (c *ServicePriceController) CreateServicePrice(ctx *fiber.Ctx) error {
	var servicePrice models.ServicePrice
	if err := ctx.BodyParser(&servicePrice); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}
	if err := c.servicePriceService.CreateServicePrice(&servicePrice); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(fiber.StatusCreated).JSON(servicePrice)
}

func (c *ServicePriceController) GetServicePrice(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}
	servicePrice, err := c.servicePriceService.GetServicePriceByID(uint(id))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "ServicePrice not found"})
	}
	return ctx.JSON(servicePrice)
}

func (c *ServicePriceController) GetAllServicePrices(ctx *fiber.Ctx) error {
	servicePrices, err := c.servicePriceService.GetAllServicePrices()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(servicePrices)
}

func (c *ServicePriceController) UpdateServicePrice(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}
	var servicePrice models.ServicePrice
	if err := ctx.BodyParser(&servicePrice); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}
	servicePrice.ID = uint(id)
	if err := c.servicePriceService.UpdateServicePrice(&servicePrice); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(servicePrice)
}

func (c *ServicePriceController) DeleteServicePrice(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}
	if err := c.servicePriceService.DeleteServicePrice(uint(id)); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.SendStatus(fiber.StatusNoContent)
}