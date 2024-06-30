package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"github.com/ycdzj/shuinotes/server/store"
	"github.com/ycdzj/shuinotes/server/utils"
)

type KB struct {
	ID    string `json:"id"`
	Owner string `json:"owner"`
	Title string `json:"title"`
}

func HandleListKBs(c *fiber.Ctx) error {
	user, err := getAuthorizedUser(c)
	if err != nil {
		return fiber.ErrInternalServerError
	}
	if user == nil {
		return fiber.ErrUnauthorized
	}

	kbs, err := store.ListKBs(map[string]interface{}{
		"owner_user_id": user.ID,
	})
	if err != nil {
		return fiber.ErrInternalServerError
	}

	respKBs := []*KB{}
	for _, kb := range kbs {
		respKBs = append(respKBs, &KB{
			ID:    kb.ExternalID,
			Owner: user.Name,
			Title: kb.Title,
		})
	}
	return c.JSON(respKBs)
}

func HandleCreateKB(c *fiber.Ctx) error {
	user, err := getAuthorizedUser(c)
	if err != nil {
		return fiber.ErrInternalServerError
	}
	if user == nil {
		return fiber.ErrUnauthorized
	}

	req := &KB{}
	if err := c.BodyParser(req); err != nil {
		log.Error().Err(err).Send()
		return fiber.ErrBadRequest
	}
	if req.Owner != user.Name {
		return fiber.ErrUnprocessableEntity
	}

	storedKB := &store.KB{
		ExternalID:  utils.RandSeq(8),
		OwnerUserID: user.ID,
		Title:       req.Title,
	}
	if err := store.CreateKB(storedKB); err != nil {
		return fiber.ErrInternalServerError
	}
	return c.JSON(&KB{
		ID:    storedKB.ExternalID,
		Owner: user.Name,
		Title: storedKB.Title,
	})
}

func HandleDeleteKB(c *fiber.Ctx) error {
	user, err := getAuthorizedUser(c)
	if err != nil {
		return fiber.ErrInternalServerError
	}
	if user == nil {
		return fiber.ErrUnauthorized
	}

	kbExternalID := c.Params("kb_id")
	if kbExternalID == "" {
		return fiber.ErrBadRequest
	}

	rowsAffected, err := store.DeleteKB(map[string]interface{}{
		"owner_user_id": user.ID,
		"external_id":   kbExternalID,
	})
	if err != nil {
		return fiber.ErrInternalServerError
	}
	if rowsAffected == 0 {
		return fiber.ErrNotFound
	}
	return nil
}

func HandleUpdateKB(c *fiber.Ctx) error {
	user, err := getAuthorizedUser(c)
	if err != nil {
		return fiber.ErrInternalServerError
	}
	if user == nil {
		return fiber.ErrUnauthorized
	}

	kbExternalID := c.Params("kb_id")
	if kbExternalID == "" {
		return fiber.ErrBadRequest
	}

	req := &KB{}
	if err := c.BodyParser(req); err != nil {
		log.Error().Err(err).Send()
		return fiber.ErrBadRequest
	}

	if req.ID != "" && req.ID != kbExternalID {
		// Modifying ID not supported
		return fiber.ErrUnprocessableEntity
	}
	if req.Owner != "" && req.Owner != user.Name {
		// Modifying owner not supported
		return fiber.ErrUnprocessableEntity
	}
	if req.Title == "" {
		return fiber.ErrUnprocessableEntity
	}

	conds := map[string]interface{}{
		"owner_user_id": user.ID,
		"external_id":   kbExternalID,
	}
	fields := map[string]interface{}{
		"title": req.Title,
	}
	rowsAffected, err := store.UpdateKB(conds, fields)
	if err != nil {
		return fiber.ErrInternalServerError
	}
	if rowsAffected == 0 {
		return fiber.ErrNotFound
	}
	return nil
}

func HandleGetKB(c *fiber.Ctx) error {
	user, err := getAuthorizedUser(c)
	if err != nil {
		return fiber.ErrInternalServerError
	}
	if user == nil {
		return fiber.ErrUnauthorized
	}

	kbExternalID := c.Params("kb_id")
	if kbExternalID == "" {
		return fiber.ErrBadRequest
	}

	kbs, err := store.ListKBs(map[string]interface{}{
		"owner_user_id": user.ID,
		"external_id":   kbExternalID,
	})
	if err != nil {
		return fiber.ErrInternalServerError
	}

	if len(kbs) == 0 {
		return fiber.ErrNotFound
	}

	kb := kbs[0]
	return c.JSON(&KB{
		ID:    kb.ExternalID,
		Owner: user.Name,
		Title: kb.Title,
	})
}
