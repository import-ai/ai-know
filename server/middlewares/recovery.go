package middlewares

import (
	"runtime/debug"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/rs/zerolog/log"
)

func NewRecovery() fiber.Handler {
	config := recover.Config{
		EnableStackTrace: true,
		StackTraceHandler: func(c *fiber.Ctx, e interface{}) {
			log.Error().
				Interface("error", e).
				Bytes("stack", debug.Stack()).
				Msg("Panic recovered")
			debug.PrintStack()
		},
	}
	return recover.New(config)
}
