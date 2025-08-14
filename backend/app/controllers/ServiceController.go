package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"app/models"
	"app/service"
)

type ServiceController struct {
	serviceService *service.ServiceService
}

func NewServiceController(serviceService *service.ServiceService) *ServiceController {
	return &ServiceController{serviceService: serviceService}
}

func (c *ServiceController) CreateService(ctx *fiber.Ctx) error {
	var service models.Service
	if err := ctx.BodyParser(&service); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}
	if err := c.serviceService.CreateService(&service); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(fiber.StatusCreated).JSON(service)
}

func (c *ServiceController) GetService(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}
	service, err := c.serviceService.GetServiceByID(uint(id))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Service not found"})
	}
	return ctx.JSON(service)
}

func (c *ServiceController) GetAllServices(ctx *fiber.Ctx) error {
	services, err := c.serviceService.GetAllServices()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(services)
}

func (c *ServiceController) UpdateService(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}
	var service models.Service
	if err := ctx.BodyParser(&service); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}
	service.ID = uint(id)
	if err := c.serviceService.UpdateService(&service); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(service)
}

func (c *ServiceController) DeleteService(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}
	if err := c.serviceService.DeleteService(uint(id)); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.SendStatus(fiber.StatusNoContent)
}