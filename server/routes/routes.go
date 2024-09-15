package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	_ "github.com/import-ai/ai-know/server/docs"
	"github.com/import-ai/ai-know/server/handlers"
	"github.com/import-ai/ai-know/server/middlewares"
)

func registerSidebarAPI(router fiber.Router) {
	router.Route("/entries", func(router fiber.Router) {
		router.Post("", handlers.CreateEntry)
		router.Route("/:entry_id", func(router fiber.Router) {
			router.Get("", handlers.GetEntry)
			router.Put("", handlers.PutEntry)
			router.Delete("", handlers.DeleteEntry)
			router.Get("/sub_entries", handlers.GetSubEntries)
			router.Post("/duplicate", handlers.DuplicateEntry)
		})
	})
}

func RegisterRoutes(router fiber.Router) {
	router.Use(middlewares.NewRecovery())

	router.Get("/swagger/*", swagger.HandlerDefault)
	router.Route("/api", func(router fiber.Router) {
		router.Route("/sidebar", func(router fiber.Router) {
			registerSidebarAPI(router)
		})
		router.Route("/workspace", func(router fiber.Router) {
			router.Get("", handlers.GetWorkspace)
		})
	})
}
