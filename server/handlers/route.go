package handlers

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(app *fiber.App) {
	app.Get("/hello", HandleHello)
	app.Post("/api/get_all_notes", HandleGetAllNotes)
	app.Post("/api/get_note_by_id", HandleGetNoteByID)
	app.Post("/api/create_note", HandleCreateNote)
	app.Post("/api/update_note", HandleUpdateNote)
}
