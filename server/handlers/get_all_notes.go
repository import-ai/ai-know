package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"github.com/ycdzj/shuinotes/server/store"
)

type Note struct {
	ID      int64  `json:"id"`
	Content string `json:"content"`
}

type GetAllNotesReply struct {
	Notes []*Note `json:"notes"`
}

func HandleGetAllNotes(c *fiber.Ctx) error {
	notes, err := store.GetAllNotes()
	if err != nil {
		log.Error().Err(err).Send()
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	reply := &GetAllNotesReply{}
	for _, note := range notes {
		reply.Notes = append(reply.Notes, &Note{
			ID:      note.ID,
			Content: note.Content,
		})
	}

	return c.JSON(reply)
}
