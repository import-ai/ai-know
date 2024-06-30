package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"github.com/ycdzj/shuinotes/server/store"
	"github.com/ycdzj/shuinotes/server/utils"
)

type Note struct {
	ID       string `json:"id"`
	KBID     string `json:"kb_id"`
	ParentID string `json:"parent_id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
}

func HandleListNotes(c *fiber.Ctx) error {
	user, err := getAuthorizedUser(c)
	if err != nil {
		return err
	}

	kbExternalID := c.Params("kb_id")
	if kbExternalID == "" {
		return fiber.ErrBadRequest
	}

	kb, err := store.GetKB(map[string]interface{}{
		"owner_user_id": user.ID,
		"external_id":   kbExternalID,
	})
	if err != nil {
		return fiber.ErrInternalServerError
	}
	if kb == nil {
		return fiber.ErrNotFound
	}

	notes, err := store.ListNotes(map[string]interface{}{
		"kb_id":          kb.ID,
		"parent_note_id": 0,
	})
	if err != nil {
		return fiber.ErrInternalServerError
	}

	respNotes := []*Note{}
	for _, note := range notes {
		respNotes = append(respNotes, &Note{
			ID:      note.ExternalID,
			KBID:    kb.ExternalID,
			Title:   note.Title,
			Content: note.Content,
		})
	}
	return c.JSON(respNotes)
}

func HandleGetNote(c *fiber.Ctx) error {
	user, err := getAuthorizedUser(c)
	if err != nil {
		return err
	}

	kbExternalID := c.Params("kb_id")
	if kbExternalID == "" {
		return fiber.ErrBadRequest
	}
	noteExternalID := c.Params("note_id")
	if noteExternalID == "" {
		return fiber.ErrBadRequest
	}

	kb, err := store.GetKB(map[string]interface{}{
		"owner_user_id": user.ID,
		"external_id":   kbExternalID,
	})
	if err != nil {
		return fiber.ErrInternalServerError
	}
	if kb == nil {
		return fiber.ErrNotFound
	}

	note, err := store.GetNote(map[string]interface{}{
		"kb_id":       kb.ID,
		"external_id": noteExternalID,
	})
	if err != nil {
		return fiber.ErrInternalServerError
	}
	if note == nil {
		return fiber.ErrNotFound
	}
	return c.JSON(&Note{
		ID:      note.ExternalID,
		KBID:    kb.ExternalID,
		Title:   note.Title,
		Content: note.Content,
	})
}

func HandleCreateNote(c *fiber.Ctx) error {
	user, err := getAuthorizedUser(c)
	if err != nil {
		return err
	}

	kbExternalID := c.Params("kb_id")
	if kbExternalID == "" {
		return fiber.ErrBadRequest
	}
	req := &Note{}
	if err := c.BodyParser(req); err != nil {
		log.Error().Err(err).Send()
		return fiber.ErrBadRequest
	}
	if req.KBID != "" && req.KBID != kbExternalID {
		return fiber.ErrUnprocessableEntity
	}

	kb, err := store.GetKB(map[string]interface{}{
		"owner_user_id": user.ID,
		"external_id":   kbExternalID,
	})
	if err != nil {
		return fiber.ErrInternalServerError
	}
	if kb == nil {
		return fiber.ErrNotFound
	}

	note := &store.Note{
		ExternalID: utils.RandSeq(8),
		KBID:       kb.ID,
		Title:      req.Title,
		Content:    req.Content,
	}
	if req.ParentID != "" {
		parentNote, err := store.GetNote(map[string]interface{}{
			"kb_id":       kb.ID,
			"external_id": req.ParentID,
		})
		if err != nil {
			return fiber.ErrInternalServerError
		}
		if parentNote == nil {
			return fiber.ErrUnprocessableEntity
		}
		note.ParentNoteID = parentNote.ID
	}

	if err := store.CreateNote(note); err != nil {
		return fiber.ErrInternalServerError
	}
	return c.JSON(&Note{
		ID: note.ExternalID,
	})
}

func HandleDeleteNote(c *fiber.Ctx) error {
	user, err := getAuthorizedUser(c)
	if err != nil {
		return err
	}

	kbExternalID := c.Params("kb_id")
	if kbExternalID == "" {
		return fiber.ErrBadRequest
	}
	noteExternalID := c.Params("note_id")
	if noteExternalID == "" {
		return fiber.ErrBadRequest
	}

	kb, err := store.GetKB(map[string]interface{}{
		"owner_user_id": user.ID,
		"external_id":   kbExternalID,
	})
	if err != nil {
		return fiber.ErrInternalServerError
	}
	if kb == nil {
		return fiber.ErrNotFound
	}

	rowsAffected, err := store.DeleteNote(map[string]interface{}{
		"kb_id":       kb.ID,
		"external_id": noteExternalID,
	})
	if err != nil {
		return fiber.ErrInternalServerError
	}
	if rowsAffected == 0 {
		return fiber.ErrNotFound
	}
	return nil
}
func HandleUpdateNote(c *fiber.Ctx) error {
	user, err := getAuthorizedUser(c)
	if err != nil {
		return err
	}

	kbExternalID := c.Params("kb_id")
	if kbExternalID == "" {
		return fiber.ErrBadRequest
	}
	noteExternalID := c.Params("note_id")
	if noteExternalID == "" {
		return fiber.ErrBadRequest
	}
	req := &Note{}
	if err := c.BodyParser(req); err != nil {
		log.Error().Err(err).Send()
		return fiber.ErrBadRequest
	}
	if req.ID != "" && req.ID != noteExternalID {
		return fiber.ErrUnprocessableEntity
	}
	if req.KBID != "" && req.KBID != kbExternalID {
		return fiber.ErrUnprocessableEntity
	}

	kb, err := store.GetKB(map[string]interface{}{
		"owner_user_id": user.ID,
		"external_id":   kbExternalID,
	})
	if err != nil {
		return fiber.ErrInternalServerError
	}
	if kb == nil {
		return fiber.ErrNotFound
	}

	conds := map[string]interface{}{
		"kb_id":       kb.ID,
		"external_id": noteExternalID,
	}
	fields := map[string]interface{}{
		"title":   req.Title,
		"content": req.Content,
	}
	rowsAffected, err := store.UpdateNote(conds, fields)
	if err != nil {
		return fiber.ErrInternalServerError
	}
	if rowsAffected == 0 {
		return fiber.ErrNotFound
	}
	return nil
}
