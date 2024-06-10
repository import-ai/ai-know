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

	listenAddr := os.Getenv("API_LISTEN_ADDR")
	if listenAddr == "" {
		log.Fatal().Msg("API_LISTEN_ADDR is empty")
	}

	if err := store.InitDB(); err != nil {
		log.Fatal().Err(err).Msg("Init DB failed")
	}
	if err := store.AutoMigrate(); err != nil {
		log.Fatal().Err(err).Msg("AutoMigrate DB failed")
	}

	app := fiber.New()
	app.Get("/hello", handlers.HandleHello)
	app.Get("/api/get_all_notes", handlers.HandleGetAllNotes)
	app.Post("/api/create_note", handlers.HandleCreateNote)
	app.Post("/api/update_note", handlers.HandleUpdateNote)
	if err := app.Listen(listenAddr); err != nil {
		log.Fatal().Err(err).Send()
	}
}
