package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/import-ai/ai-know/server/db"
)

type Entries struct {
	Private string `json:"private" example:"1000001"`
	Team    string `json:"team" example:"1000002"`
}

// GetWorkspace
//
//	@Summary		Get Workspace
//	@Description	Get properties of current workspace.
//	@Tags			Workspace
//	@Router			/api/workspace [get]
//	@Success		200	{object}	handlers.GetWorkspace.Resp
func GetWorkspace(c *fiber.Ctx) error {
	type Resp struct {
		Entries *Entries `json:"entries"`
	}
	ctx := c.Context()
	workspace, err := db.GetWorkspace(ctx, 1)
	if err != nil {
		return err
	}
	if workspace == nil {
		workspace, err = db.CreateFirstWorkspace(ctx)
		if err != nil {
			return err
		}
	}
	return c.JSON(&Resp{
		Entries: &Entries{
			Private: strconv.FormatInt(workspace.PrivateSidebarEntry, 10),
			Team:    strconv.FormatInt(workspace.TeamSidebarEntry, 10),
		},
	})
}
