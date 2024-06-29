package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"github.com/ycdzj/shuinotes/server/store"
)

type User struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

func HandleGetUser(c *fiber.Ctx) error {
	return nil
}

func HandleCreateUser(c *fiber.Ctx) error {
	user := &User{}
	if err := c.BodyParser(user); err != nil {
		log.Error().Err(err).Send()
		return fiber.ErrBadRequest
	}
	if user.User == "" || user.Password == "" {
		return fiber.ErrBadRequest
	}

	if dbUser, err := store.GetUserByName(user.User); err != nil {
		log.Error().Err(err).Send()
		return fiber.ErrInternalServerError
	} else if dbUser != nil {
		return fiber.ErrUnprocessableEntity
	}

	if err := store.CreateUser(&store.User{
		UniqueName:   user.User,
		PasswordHash: getPasswordHash(user.Password),
	}); err != nil {
		log.Error().Err(err).Send()
		return fiber.ErrInternalServerError
	}

	return nil
}
