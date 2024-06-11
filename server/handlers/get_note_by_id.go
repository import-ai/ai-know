package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"github.com/ycdzj/shuinotes/server/store"
)

type GetNoteByIDRequest struct {
	ID int64 `json:"id"`
}
type GetNoteByIDReply struct {
	Note *Note `json:"note"`
}

func HandleGetNoteByID(c *fiber.Ctx) error {
	req := &GetNoteByIDRequest{}
	if err := c.BodyParser(req); err != nil {
		log.Error().Err(err).Send()
		return err
	}

	note, err := store.GetNoteByID(req.ID)
	if err != nil {
		log.Error().Err(err).Send()
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	if note == nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	reply := &GetNoteByIDReply{
		Note: &Note{
			ID:      note.ID,
			Content: note.Content,
		},
	}
	return c.JSON(reply)
}
