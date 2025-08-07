package controllers

import "github.com/gofiber/fiber/v2"

func Ok(c *fiber.Ctx) error {
	c.Status(fiber.StatusOK)
	return nil
}