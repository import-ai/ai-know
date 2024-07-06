package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"github.com/ycdzj/shuinotes/server/auth"
	"github.com/ycdzj/shuinotes/server/config"
	"github.com/ycdzj/shuinotes/server/store"
	"github.com/ycdzj/shuinotes/server/utils"
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
