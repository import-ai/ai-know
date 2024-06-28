package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/rs/zerolog/log"
)

func RegisterRecover(router fiber.Router) {
	config := recover.Config{
		EnableStackTrace: true,
		StackTraceHandler: func(c *fiber.Ctx, e interface{}) {
			log.Error().Interface("error", e).Msg("Panic recovered")
		},
	}
	router.Use(recover.New(config))
}
