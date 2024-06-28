package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/rs/zerolog/log"
)

func NewRecover() fiber.Handler {
	config := recover.Config{
		EnableStackTrace: true,
		StackTraceHandler: func(c *fiber.Ctx, e interface{}) {
			log.Error().Interface("error", e).Msg("Panic recovered")
		},
	}
	return recover.New(config)
}
