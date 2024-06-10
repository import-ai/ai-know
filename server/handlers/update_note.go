package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"github.com/ycdzj/shuinotes/server/store"
)

type UpdateNoteRequest struct {
	ID      int64  `json:"id"`
	Content string `json:"content"`
}

func HandleUpdateNote(c *fiber.Ctx) error {
	req := &UpdateNoteRequest{}
	if err := c.BodyParser(req); err != nil {
		log.Error().Err(err).Send()
		return err
	}

	if err := store.UpdateNoteContent(req.ID, req.Content); err != nil {
		log.Error().Err(err).Send()
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendStatus(fiber.StatusOK)
}
