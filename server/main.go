package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/ycdzj/shuinotes/server/config"
	"github.com/ycdzj/shuinotes/server/handlers"
	"github.com/ycdzj/shuinotes/server/store"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	config.InitFromEnv()

	if err := store.InitDB(); err != nil {
		log.Fatal().Err(err).Msg("Init DB failed")
	}
	if err := store.AutoMigrate(); err != nil {
		log.Fatal().Err(err).Msg("AutoMigrate DB failed")
	}

	app := fiber.New()
	handlers.RegisterRoutes(app)
	if err := app.Listen(config.ListenAddr()); err != nil {
		log.Fatal().Err(err).Send()
	}
}

// Knowledge Base
// /kbs/{kb_id}
// /kbs/{kb_id}/notes
// /kbs/{kb_id}/notes/{note_id}
// Note
