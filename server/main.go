package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/import-ai/ai-know/server/config"
	"github.com/import-ai/ai-know/server/routes"
	"github.com/import-ai/ai-know/server/store"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Logger = log.With().Caller().Logger()
	config.InitFromEnv()

	if err := store.InitDB(config.DataSourceName()); err != nil {
		log.Fatal().Err(err).Msg("Init DB failed")
	}
	if err := store.AutoMigrate(); err != nil {
		log.Fatal().Err(err).Msg("AutoMigrate DB failed")
	}

	app := fiber.New()
	routes.RegisterRoutes(app)
	if err := app.Listen(config.ListenAddr()); err != nil {
		log.Fatal().Err(err).Send()
	}
}
