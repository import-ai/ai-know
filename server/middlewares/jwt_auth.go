package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ycdzj/shuinotes/server/auth"
	"github.com/ycdzj/shuinotes/server/config"
)

func NewJWTAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenStr := c.Cookies(config.JWTCookieName())
		user, ok := auth.ValidateJWT(tokenStr)
		if !ok {
			return fiber.ErrUnauthorized
		}
		c.Locals("user", user)
		return c.Next()
	}
}
