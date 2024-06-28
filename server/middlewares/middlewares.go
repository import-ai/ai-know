package middlewares

import "github.com/gofiber/fiber/v2"

func RegisterMiddlewares(router fiber.Router) {
	RegisterRecover(router)
	RegisterAuth(router)
}
