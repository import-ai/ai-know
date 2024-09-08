package handlers

import "github.com/gofiber/fiber/v2"

type Entry struct {
	ID            string `json:"id" example:"1000005"`
	Title         string `json:"title" example:"Note Title"`
	Type          string `json:"type" example:"note" enums:"note,group,link"`
	HasSubEntries bool   `json:"has_sub_entries" example:"false"`
}

// CreateEntry
//
//	@Summary		Create Entry
//	@Description	Create an entry with the specified properties.
//	@Description	|      Field      | Required |      Description      |
//	@Description	| :-------------: | :------: | :-------------------: |
//	@Description	|      title      |   Yes    |  Title of new entry   |
//	@Description	|      type       |   Yes    |   Type of new entry   |
//	@Description	|     parent      |   Yes    |    Parent entry ID    |
//	@Description	| posistion_after |    No    | Position of new entry |
//	@Description
//	@Description	If `position_after` is empty, the new entry will be the first in parent's sub-entries. Otherwise, it's positioned after the specified sub-entry.
//	@Tags			Sidebar
//	@Router			/api/sidebar/entries [post]
//	@Param			Body	body		handlers.CreateEntry.Req	true	"Request Body"
//	@Success		200		{object}	handlers.CreateEntry.Resp
func CreateEntry(c *fiber.Ctx) error {
	type Req struct {
		Title         string `json:"title" example:"Note Title"`
		Type          string `json:"type" example:"note"`
		Parent        string `json:"parent" example:"10000003"`
		PositionAfter string `json:"position_after" example:"10000002"`
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
//	@Router			/api/sidebar/entries/{entry_id} [get]
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
//	@Description	|      Field      | Required |      Description      |
//	@Description	| :-------------: | :------: | :-------------------: |
//	@Description	|      title      |    No    |  Title of the entry   |
//	@Description	|     parent      |    No    |    Parent entry ID    |
//	@Description	| posistion_after |    No    | Position of the entry |
//	@Description
//	@Description	> - If `title` is non-empty, update title of the entry.
//	@Description	> - If `parent` is non-empty, move the entry to the specified parent entry.
//	@Description	>     - If `position_after` is empty, the new entry will be the first in parent's sub-entries. Otherwise, it's positioned after the specified sub-entry.
//	@Tags			Sidebar
//	@Router			/api/sidebar/entries/{entry_id} [put]
//	@Param			entry_id	path		string					true	"Entry ID"
//	@Param			Body		body		handlers.PutEntry.Req	true	"Request Body"
//	@Success		200			{object}	handlers.PutEntry.Resp
func PutEntry(c *fiber.Ctx) error {
	type Req struct {
		Title         string `json:"title" example:"Note Title"`
		Parent        string `json:"parent" example:"10000003"`
		PositionAfter string `json:"position_after" example:"10000002"`
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
//	@Router			/api/sidebar/entries/{entry_id} [delete]
func DeleteEntry(c *fiber.Ctx) error {
	return nil
}

// GetSubEntries
//
//	@Summary		Get Sub-Entries
//	@Description	Get sub-entries of an entry.
//	@Tags			Sidebar
//	@Router			/api/sidebar/entries/{entry_id}/sub_entries [get]
//	@Param			entry_id	path		string	true	"Entry ID"
//	@Success		200			{object}	handlers.GetSubEntries.Resp
func GetSubEntries(c *fiber.Ctx) error {
	type Resp struct {
		SubEntries []*Entry `json:"sub_entries"`
	}
	return nil
}

// DuplicateEntry
//
//	@Summary		Duplicate Entry
//	@Description	Duplicate an entry.
//	@Description	|      Field      | Required |      Description      |
//	@Description	| :-------------: | :------: | :-------------------: |
//	@Description	|      title      |    No    |  Title of new entry   |
//	@Description	|     parent      |   Yes    |    Parent entry ID    |
//	@Description	| posistion_after |    No    | Position of new entry |
//	@Description
//	@Description	If `title` is empty, it will default to the old entry’s title.
//	@Description	If `position_after` is empty, the new entry will be the first in parent's sub-entries. Otherwise, it's positioned after the specified sub-entry.
//	@Tags			Sidebar
//	@Router			/api/sidebar/entries/{entry_id}/duplicate [post]
//	@Param			entry_id	path		string						true	"Entry ID"
//	@Param			Body		body		handlers.DuplicateEntry.Req	true	"Request Body"
//	@Success		200			{object}	handlers.DuplicateEntry.Resp
func DuplicateEntry(c *fiber.Ctx) error {
	type Req struct {
		Title         string `json:"title" example:"Note Title"`
		Parent        string `json:"parent" example:"10000003"`
		PositionAfter string `json:"position_after" example:"10000002"`
	}
	type Resp struct {
		Entry *Entry `json:"entry"`
	}
	return nil
}
