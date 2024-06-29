package middlewares

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
)

type JWTAuthConfig struct {
	SecretKey  string
	CookieName string
}

func NewJWTAuth(cfg *JWTAuthConfig) fiber.Handler {
	return func(c *fiber.Ctx) error {
		tokenStr := c.Cookies(cfg.CookieName)
		if tokenStr == "" {
			return fiber.ErrUnauthorized
		}

		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			return cfg.SecretKey, nil
		})
		if err != nil {
			log.Error().Err(err).Send()
			return fiber.ErrUnauthorized
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return fiber.ErrUnauthorized
		}

		exp, err := claims.GetExpirationTime()
		if !exp.Before(time.Now()) {
			return fiber.ErrUnauthorized
		}

		user, ok := claims["user"].(string)
		if !ok {
			return fiber.ErrUnauthorized
		}

		c.Locals("user", user)
		return c.Next()
	}
}
