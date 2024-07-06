package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ycdzj/shuinotes/server/handlers"
	"github.com/ycdzj/shuinotes/server/middlewares"
)

func RegisterRoutes(router fiber.Router) {
	router.Use(middlewares.NewRecovery())

	router = router.Group("/api")

	router.Post("/login", handlers.HandleLogin)
	router.Post("/logout", handlers.HandleLogout)
	router.Post("/register", handlers.HandleRegister)

	router.Use(middlewares.NewJWTAuth())

	prefix := router.Group("/user")
	prefix.Get("", handlers.HandleGetAuthorizedUser)

	prefix = router.Group("/kbs")
	prefix.Get("", handlers.HandleListKBs)
	prefix.Post("", handlers.HandleCreateKB)

	prefix = router.Group("/kbs/:kb_id")
	prefix.Get("", handlers.HandleGetKB)
	prefix.Put("", handlers.HandleUpdateKB)
	prefix.Delete("", handlers.HandleDeleteKB)

	prefix = router.Group("/kbs/:kb_id/notes")
	prefix.Get("", handlers.HandleListNotes)
	prefix.Post("", handlers.HandleCreateNote)

	prefix = router.Group("/kbs/:kb_id/notes/:note_id")
	prefix.Get("", handlers.HandleGetNote)
	prefix.Put("", handlers.HandleUpdateNote)
	prefix.Delete("", handlers.HandleDeleteNote)
}
