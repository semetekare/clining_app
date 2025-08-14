package controllers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"clining_app/app/models"
	"clining_app/app/service"
)

type RoomTypeController struct {
	roomTypeService *service.RoomTypeService
}

func NewRoomTypeController(roomTypeService *service.RoomTypeService) *RoomTypeController {
	return &RoomTypeController{roomTypeService: roomTypeService}
}

func (c *RoomTypeController) CreateRoomType(ctx *fiber.Ctx) error {
	var roomType models.RoomType
	if err := ctx.BodyParser(&roomType); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}
	if err := c.roomTypeService.CreateRoomType(&roomType); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(fiber.StatusCreated).JSON(roomType)
}

func (c *RoomTypeController) GetRoomType(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}
	roomType, err := c.roomTypeService.GetRoomTypeByID(uint(id))
	if err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "RoomType not found"})
	}
	return ctx.JSON(roomType)
}

func (c *RoomTypeController) GetAllRoomTypes(ctx *fiber.Ctx) error {
	roomTypes, err := c.roomTypeService.GetAllRoomTypes()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(roomTypes)
}

func (c *RoomTypeController) UpdateRoomType(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}
	var roomType models.RoomType
	if err := ctx.BodyParser(&roomType); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}
	roomType.ID = uint(id)
	if err := c.roomTypeService.UpdateRoomType(&roomType); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.JSON(roomType)
}

func (c *RoomTypeController) DeleteRoomType(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}
	if err := c.roomTypeService.DeleteRoomType(uint(id)); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.SendStatus(fiber.StatusNoContent)
}