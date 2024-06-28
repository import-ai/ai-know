package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func NewJWTAuth(jwtSecretKey string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		log.Info().Msg("JWTAuth called")
		return c.Next()
	}
}
