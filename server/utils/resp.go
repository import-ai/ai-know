package utils

import (
	"github.com/gofiber/fiber/v2"
)

type ErrorResp struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func MakeErrorResp(c *fiber.Ctx, status int, msg string) error {
	return c.Status(status).JSON(&ErrorResp{
		Status:  status,
		Message: msg,
	})
}

func MakeOKResp(c *fiber.Ctx, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(data)
}
