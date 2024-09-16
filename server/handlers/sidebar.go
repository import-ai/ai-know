package handlers

import (
	"errors"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/import-ai/ai-know/server/db"
	"github.com/import-ai/ai-know/server/sql/queries"
)

type Entry struct {
	ID            string `json:"id" example:"1000005"`
	Title         string `json:"title" example:"Note Title"`
	Type          string `json:"type" example:"note" enums:"note,group,link"`
	HasSubEntries bool   `json:"has_sub_entries" example:"false"`
}

var ErrInvalidEntryType = errors.New("Invalid entry type")
var ErrEntryNotExist = errors.New("Entry not exist")

func isValidEntryType(t string) bool {
	validTypes := []queries.SidebarEntryType{
		queries.SidebarEntryTypeLink,
		queries.SidebarEntryTypeNote,
		queries.SidebarEntryTypeGroup,
	}
	for _, validType := range validTypes {
		if string(validType) == t {
			return true
		}
	}
	return false
}

const kEntryID = "entry_id"

func parsePosition(parent string, positionAfter string) (int64, int64, error) {
	parentID, err := strconv.ParseInt(parent, 10, 64)
	if err != nil {
		return 0, 0, err
	}
	if positionAfter == "" {
		return parentID, parentID, nil
	}
	prevID, err := strconv.ParseInt(positionAfter, 10, 64)
	if err != nil {
		return 0, 0, err
	}
	return parentID, prevID, nil
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

	ctx := c.Context()

	var req *Req
	if err := c.BodyParser(&req); err != nil {
		return err
	}
	if !isValidEntryType(req.Type) {
		return ErrInvalidEntryType
	}

	parentID, prevID, err := parsePosition(req.Parent, req.PositionAfter)
	if err != nil {
		return err
	}

	entry, err := db.CreateSidebarEntry(ctx, &db.CreateSidebarEntryArgs{
		Title:  req.Title,
		Type:   queries.SidebarEntryType(req.Type),
		Parent: parentID,
		PrevID: prevID,
	})
	if err != nil {
		return err
	}

	return c.JSON(&Resp{
		Entry: &Entry{
			ID:    strconv.FormatInt(entry.ID, 10),
			Title: entry.Title,
			Type:  string(entry.Type),
		},
	})
}

func ParseEntryID(c *fiber.Ctx) error {
	entryID, err := strconv.ParseInt(c.Params("entry_id"), 10, 64)
	if err != nil {
		return err
	}
	c.Locals(kEntryID, entryID)
	return c.Next()
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

	ctx := c.Context()
	entryID := c.Locals(kEntryID).(int64)

	entry, err := db.GetSidebarEntry(ctx, entryID)
	if err != nil {
		return err
	}
	if entry == nil {
		return fiber.ErrNotFound
	}

	subEntry, err := db.GetSidebarSubEntry(ctx, entry.ID, entry.ID)
	if err != nil {
		return err
	}

	return c.JSON(&Resp{
		Entry: &Entry{
			ID:            strconv.FormatInt(entry.ID, 10),
			Title:         entry.Title,
			Type:          string(entry.Type),
			HasSubEntries: subEntry != nil,
		},
	})
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

	ctx := c.Context()
	entryID := c.Locals(kEntryID).(int64)

	var req *Req
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	args := &db.PutSidebarEntryArgs{
		EntryID: entryID,
		Title:   req.Title,
	}

	if req.Parent != "" {
		var err error
		args.Parent, err = strconv.ParseInt(req.Parent, 10, 64)
		if err != nil {
			return err
		}
	}
	if req.PositionAfter != "" {
		var err error
		args.PrevID, err = strconv.ParseInt(req.PositionAfter, 10, 64)
		if err != nil {
			return err
		}
	}
	if args.PrevID == 0 {
		args.PrevID = args.Parent
	}

	return db.PutSidebarEntry(ctx, args)
}

// DeleteEntry
//
//	@Summary		Delete Entry
//	@Description	Delete an entry and all its sub-entries.
//	@Tags			Sidebar
//	@Param			entry_id	path	string	true	"Entry ID"
//	@Router			/api/sidebar/entries/{entry_id} [delete]
func DeleteEntry(c *fiber.Ctx) error {
	ctx := c.Context()
	entryID := c.Locals(kEntryID).(int64)
	return db.RemoveSidebarEntry(ctx, entryID)
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

	ctx := c.Context()
	parentID := c.Locals(kEntryID).(int64)
	prevID := parentID
	entries := []*queries.SidebarEntry{}
	for {
		entry, err := db.GetSidebarSubEntry(ctx, parentID, prevID)
		if err != nil {
			return err
		}
		if entry == nil {
			break
		}
		entries = append(entries, entry)
		prevID = entry.ID
	}

	respEntries := []*Entry{}
	for _, entry := range entries {
		subEntry, err := db.GetSidebarSubEntry(ctx, entry.ID, entry.ID)
		if err != nil {
			return err
		}
		respEntries = append(respEntries, &Entry{
			ID:            strconv.FormatInt(entry.ID, 10),
			Title:         entry.Title,
			Type:          string(entry.Type),
			HasSubEntries: subEntry != nil,
		})
	}

	return c.JSON(&Resp{
		SubEntries: respEntries,
	})
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

	ctx := c.Context()
	entryID := c.Locals(kEntryID).(int64)

	var req *Req
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	parentID, prevID, err := parsePosition(req.Parent, req.PositionAfter)
	if err != nil {
		return err
	}

	args := &db.CreateSidebarEntryArgs{
		Parent: parentID,
		PrevID: prevID,
	}

	entry, err := db.GetSidebarEntry(ctx, entryID)
	if err != nil {
		return err
	}
	if entry == nil {
		return ErrEntryNotExist
	}

	args.Type = entry.Type
	if req.Title != "" {
		args.Title = req.Title
	} else {
		args.Title = entry.Title
	}

	newEntry, err := db.CreateSidebarEntry(c.Context(), args)
	if err != nil {
		return err
	}

	return c.JSON(&Resp{
		Entry: &Entry{
			ID:            strconv.FormatInt(newEntry.ID, 10),
			Title:         newEntry.Title,
			Type:          string(newEntry.Type),
			HasSubEntries: false,
		},
	})
}
