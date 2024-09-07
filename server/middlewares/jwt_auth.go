package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/import-ai/ai-know/server/auth"
	"github.com/import-ai/ai-know/server/config"
	"github.com/import-ai/ai-know/server/store"
	"github.com/import-ai/ai-know/server/utils"
	"github.com/rs/zerolog/log"
)

func NewJWTAuth() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenStr := c.Cookies(config.JWTCookieName())
		userName, ok := auth.ValidateJWT(tokenStr)
		if !ok {
			return utils.MakeErrorResp(c, fiber.StatusUnauthorized, "Not logged in")
		}
		user, err := store.GetUserByName(userName)
		if err != nil {
			log.Error().Err(err).Send()
			return utils.MakeErrorResp(c, fiber.StatusInternalServerError, "")
		}
		if user == nil {
			return utils.MakeErrorResp(c, fiber.StatusUnauthorized, "Not logged in")
		}
		c.Locals("authorized_user", user)
		return c.Next()
	}
}
