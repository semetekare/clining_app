package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"clining_app/app/models"
	"clining_app/app/service"
)

type SurchargeController struct {
	surchargeService *service.SurchargeService
}

func NewSurchargeController(surchargeService *service.SurchargeService) *SurchargeController {
	return &SurchargeController{surchargeService: surchargeService}
}

func (c *SurchargeController) CreateSurcharge(ctx *fiber.Ctx) error {
	var surcharge models.Surcharge
	if err := ctx.BodyParser(&surcharge); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}
	if err := c.surchargeService.CreateSurcharge(&surcharge); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(fiber.StatusCreated).JSON(surcharge)
}

func (c *SurchargeController) GetSurcharge(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}
	surcharge, err := c.surchargeService.GetSurchargeByID(uint(id))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Surcharge not found"})
	}
	return ctx.JSON(surcharge)
}

func (c *SurchargeController) GetAllSurcharges(ctx *fiber.Ctx) error {
	surcharges, err := c.surchargeService.GetAllSurcharges()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(surcharges)
}

func (c *SurchargeController) UpdateSurcharge(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}
	var surcharge models.Surcharge
	if err := ctx.BodyParser(&surcharge); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}
	surcharge.ID = uint(id)
	if err := c.surchargeService.UpdateSurcharge(&surcharge); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(surcharge)
}

func (c *SurchargeController) DeleteSurcharge(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}
	if err := c.surchargeService.DeleteSurcharge(uint(id)); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.SendStatus(fiber.StatusNoContent)
}