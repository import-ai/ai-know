package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/import-ai/ai-know/server/store"
	"github.com/import-ai/ai-know/server/utils"
	"github.com/rs/zerolog/log"
)

type KB struct {
	ID    string `json:"id"`
	Owner string `json:"owner"`
	Title string `json:"title"`
}

func HandleListKBs(c *fiber.Ctx) error {
	user, ok := c.Locals("authorized_user").(*store.User)
	if !ok {
		return utils.MakeErrorResp(c, fiber.StatusInternalServerError, "")
	}

	kbs, err := store.ListKBs(map[string]interface{}{
		"owner_user_id": user.ID,
	})
	if err != nil {
		return utils.MakeErrorResp(c, fiber.StatusInternalServerError, "")
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
	user, ok := c.Locals("authorized_user").(*store.User)
	if !ok {
		return utils.MakeErrorResp(c, fiber.StatusInternalServerError, "")
	}

	req := &KB{}
	if err := c.BodyParser(req); err != nil {
		log.Error().Err(err).Send()
		return utils.MakeErrorResp(c, fiber.StatusBadRequest, "Not a valid json")
	}
	if req.Title == "" {
		return utils.MakeErrorResp(c, fiber.StatusBadRequest, "Empty title")
	}
	if req.Owner != "" && req.Owner != user.Name {
		return utils.MakeErrorResp(c, fiber.StatusUnprocessableEntity, "Cannot create KB owned by others")
	}

	storedKB := &store.KB{
		ExternalID:  utils.RandSeq(8),
		OwnerUserID: user.ID,
		Title:       req.Title,
	}
	if err := store.CreateKB(storedKB); err != nil {
		return utils.MakeErrorResp(c, fiber.StatusInternalServerError, "")
	}
	return c.JSON(&KB{
		ID:    storedKB.ExternalID,
		Owner: user.Name,
		Title: storedKB.Title,
	})
}

func HandleDeleteKB(c *fiber.Ctx) error {
	user, ok := c.Locals("authorized_user").(*store.User)
	if !ok {
		return utils.MakeErrorResp(c, fiber.StatusInternalServerError, "")
	}

	kbExternalID := c.Params("kb_id")
	if kbExternalID == "" {
		return utils.MakeErrorResp(c, fiber.StatusBadRequest, "Empty kb_id")
	}

	rowsAffected, err := store.DeleteKB(map[string]interface{}{
		"owner_user_id": user.ID,
		"external_id":   kbExternalID,
	})
	if err != nil {
		return utils.MakeErrorResp(c, fiber.StatusInternalServerError, "")
	}
	if rowsAffected == 0 {
		return fiber.ErrNotFound
	}
	return nil
}

func HandleUpdateKB(c *fiber.Ctx) error {
	user, ok := c.Locals("authorized_user").(*store.User)
	if !ok {
		return utils.MakeErrorResp(c, fiber.StatusInternalServerError, "")
	}

	kbExternalID := c.Params("kb_id")
	if kbExternalID == "" {
		return utils.MakeErrorResp(c, fiber.StatusBadRequest, "Empty kb_id")
	}

	req := &KB{}
	if err := c.BodyParser(req); err != nil {
		log.Error().Err(err).Send()
		return utils.MakeErrorResp(c, fiber.StatusBadRequest, "Not a valid json")
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
		return utils.MakeErrorResp(c, fiber.StatusInternalServerError, "")
	}
	if rowsAffected == 0 {
		return fiber.ErrNotFound
	}
	return nil
}

func HandleGetKB(c *fiber.Ctx) error {
	user, ok := c.Locals("authorized_user").(*store.User)
	if !ok {
		return utils.MakeErrorResp(c, fiber.StatusInternalServerError, "")
	}

	kbExternalID := c.Params("kb_id")
	if kbExternalID == "" {
		return utils.MakeErrorResp(c, fiber.StatusBadRequest, "Empty kb_id")
	}

	kbs, err := store.ListKBs(map[string]interface{}{
		"owner_user_id": user.ID,
		"external_id":   kbExternalID,
	})
	if err != nil {
		return utils.MakeErrorResp(c, fiber.StatusInternalServerError, "")
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
