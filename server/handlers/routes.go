package handlers

import "github.com/gofiber/fiber/v2"

func RegisterRoutes(router fiber.Router) {
	prefix := router.Group("/kbs")
	prefix.Get("", HandleListKBs)
	prefix.Post("", HandleCreateKB)

	prefix = router.Group("/kbs/:kb_id")
	prefix.Get("", HandleGetKB)
	prefix.Put("", HandleUpdateKB)
	prefix.Delete("", HandleDeleteKB)

	prefix = router.Group("kbs/:kb_id/notes")
	prefix.Get("", HandleListNotes)
	prefix.Post("", HandleCreateNote)

	prefix = router.Group("kbs/:kb_id/notes/:note_id")
	prefix.Get("", HandleGetNote)
	prefix.Put("", HandleUpdateNote)
	prefix.Delete("", HandleDeleteNote)
}
