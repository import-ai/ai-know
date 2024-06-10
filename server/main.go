package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/ycdzj/shuinotes/server/handlers"
	"github.com/ycdzj/shuinotes/server/store"
)

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	if err := store.InitDB(); err != nil {
		log.Fatal().Err(err).Msg("Init DB failed")
	}

	app := fiber.New()
	app.Get("/hello", handlers.HandleHello)
	app.Listen(":3456")
}
