package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"clining_app/app/models"
	"clining_app/app/service"
)

type ExtraServiceController struct {
	extraServiceService *service.ExtraServiceService
}

func NewExtraServiceController(extraServiceService *service.ExtraServiceService) *ExtraServiceController {
	return &ExtraServiceController{extraServiceService: extraServiceService}
}

func (c *ExtraServiceController) CreateExtraService(ctx *fiber.Ctx) error {
	var extraService models.ExtraService
	if err := ctx.BodyParser(&extraService); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}
	if err := c.extraServiceService.CreateExtraService(&extraService); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(fiber.StatusCreated).JSON(extraService)
}

func (c *ExtraServiceController) GetExtraService(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}
	extraService, err := c.extraServiceService.GetExtraServiceByID(uint(id))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "ExtraService not found"})
	}
	return ctx.JSON(extraService)
}

func (c *ExtraServiceController) GetAllExtraServices(ctx *fiber.Ctx) error {
	extraServices, err := c.extraServiceService.GetAllExtraServices()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(extraServices)
}

func (c *ExtraServiceController) UpdateExtraService(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}
	var extraService models.ExtraService
	if err := ctx.BodyParser(&extraService); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}
	extraService.ID = uint(id)
	if err := c.extraServiceService.UpdateExtraService(&extraService); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(extraService)
}

func (c *ExtraServiceController) DeleteExtraService(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}
	if err := c.extraServiceService.DeleteExtraService(uint(id)); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.SendStatus(fiber.StatusNoContent)
}