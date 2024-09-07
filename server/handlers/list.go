package handlers

import "github.com/gofiber/fiber/v2"

type Entry struct {
	ID            string `json:"id"`
	Title         string `json:"title"`
	Type          string `json:"type"`
	Parent        string `json:"parent"`
	PositionAfter string `json:"position_after"`
	HasSubEntries bool   `json:"has_sub_entries"`
}

// CreateEntry
//
//	@Summary		Create Entry
//	@Description	Create an entry with the specified properties.
//	@Tags			Sidebar
//	@Router			/api/sidebar/list/entries [post]
//	@Param			Body	body		handlers.CreateEntry.Req	true	"Request Body"
//	@Success		200		{object}	handlers.CreateEntry.Resp
func CreateEntry(c *fiber.Ctx) error {
	type Req struct {
		Title         string `json:"title"`
		Type          string `json:"type"`
		Parent        string `json:"parent"`
		PositionAfter string `json:"position_after"`
	}
	type Resp struct {
		Entry *Entry `json:"entry"`
	}

	return nil
}

// GetEntry
//
//	@Summary		Get Entry
//	@Description	Get properties of an entry.
//	@Tags			Sidebar
//	@Router			/api/sidebar/list/entries/{entry_id} [get]
//	@Param			entry_id	path		string	true	"Entry ID"
//	@Success		200			{object}	handlers.GetEntry.Resp
func GetEntry(c *fiber.Ctx) error {
	type Resp struct {
		Entry *Entry `json:"entry"`
	}
	return nil
}

// PutEntry
//
//	@Summary		Update Entry
//	@Description	Update properties of an entry.
//	@Tags			Sidebar
//	@Router			/api/sidebar/list/entries/{entry_id} [put]
//	@Param			Body	body		handlers.PutEntry.Req	true	"Request Body"
//	@Success		200		{object}	handlers.PutEntry.Resp
func PutEntry(c *fiber.Ctx) error {
	type Req struct {
		Title         string `json:"title"`
		Parent        string `json:"parent"`
		PositionAfter string `json:"position_after"`
	}
	type Resp struct {
		Entry *Entry `json:"entry"`
	}

	return nil
}

// DeleteEntry
//
//	@Summary		Delete Entry
//	@Description	Delete an entry and all its sub-entries.
//	@Tags			Sidebar
//	@Param			entry_id	path	string	true	"Entry ID"
//	@Router			/api/sidebar/list/entries/{entry_id} [delete]
func DeleteEntry(c *fiber.Ctx) error {
	return nil
}

// GetSubEntries
//
//	@Summary		Get Sub-Entries
//	@Description	Get sub-entries of an entry.
//	@Tags			Sidebar
//	@Param			entry_id	path		string	true	"Entry ID"
//	@Success		200			{object}	handlers.GetSubEntries.Resp
//	@Router			/api/sidebar/list/entries/{entry_id}/sub_entries [get]
func GetSubEntries(c *fiber.Ctx) error {
	type Resp struct {
		SubEntries []*Entry `json:"sub_entries"`
	}
	return nil
}

func DuplicateEntry(c *fiber.Ctx) error {
	return nil
}
