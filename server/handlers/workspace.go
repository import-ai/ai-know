package handlers

import "github.com/gofiber/fiber/v2"

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
	return nil
}
