package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func HandleHello(c *fiber.Ctx) error {
	err := c.SendString("Hello, world!")
	return err
}
