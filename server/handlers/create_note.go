package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"github.com/ycdzj/shuinotes/server/store"
)

type CreateNoteRequest struct {
	Content string `json:"content"`
}

type CreateNoteReply struct {
	ID int64 `json:"id"`
}

func HandleCreateNote(c *fiber.Ctx) error {
	req := &CreateNoteRequest{}
	if err := c.BodyParser(req); err != nil {
		log.Error().Err(err).Send()
		return err
	}

	note := &store.Note{
		Content: req.Content,
	}
	if err := store.CreateNote(note); err != nil {
		log.Error().Err(err).Send()
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	reply := &CreateNoteReply{
		ID: note.ID,
	}
	return c.JSON(reply)
}
