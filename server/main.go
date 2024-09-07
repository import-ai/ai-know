package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/import-ai/ai-know/server/config"
	"github.com/import-ai/ai-know/server/routes"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

// AIKnow API
//
//	@title			AIKnow API
//	@version		1.0
//	@license.name	GPLv3
func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Logger = log.With().Caller().Logger()
	config.InitFromEnv()

	app := fiber.New()
	routes.RegisterRoutes(app)
	if err := app.Listen(config.ListenAddr()); err != nil {
		log.Fatal().Err(err).Send()
	}
}
