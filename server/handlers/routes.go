package handlers

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(app *fiber.App) {
	app.All("/kbs", HandleKBs)
	app.All("/kbs/:kb_id", HandleSingleKB)
	app.All("/kbs/:kb_id/notes", HandleNotes)
	app.All("/kbs/:kb_id/notes/:note_id", HandleSingleNote)
}
