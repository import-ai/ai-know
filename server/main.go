package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/ycdzj/shuinotes/server/config"
	"github.com/ycdzj/shuinotes/server/routes"
	"github.com/ycdzj/shuinotes/server/store"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
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
