package handlers

import (
	"crypto/sha256"
	"encoding/hex"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"github.com/ycdzj/shuinotes/server/auth"
	"github.com/ycdzj/shuinotes/server/config"
	"github.com/ycdzj/shuinotes/server/store"
)

type User struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

func (s *User) PasswordHash() string {
	hashBytes := sha256.Sum256([]byte(s.User + "#" + s.Password))
	return hex.EncodeToString(hashBytes[:])
}

func HandleLogin(c *fiber.Ctx) error {
	req := &User{}
	if err := c.BodyParser(req); err != nil {
		log.Error().Err(err).Send()
		return fiber.ErrBadRequest
	}
	if req.User == "" || req.Password == "" {
		return fiber.ErrBadRequest
	}

	storedUser, err := store.GetUserByName(req.User)
	if err != nil {
		log.Error().Err(err).Send()
		return fiber.ErrInternalServerError
	}
	if storedUser == nil || storedUser.PasswordHash != req.PasswordHash() {
		return fiber.ErrUnprocessableEntity
	}

	tokenStr, err := auth.GenerateJWT(req.User, time.Now().Add(time.Hour*24))
	if err != nil {
		return fiber.ErrInternalServerError
	}
	c.Cookie(&fiber.Cookie{
		Name:  config.JWTCookieName(),
		Value: tokenStr,
	})
	return c.SendStatus(fiber.StatusOK)
}

func HandleGetUser(c *fiber.Ctx) error {
	authUser, ok := c.Locals("user").(string)
	if !ok {
		return fiber.ErrInternalServerError
	}

	if c.Params("user_name") != authUser {
		return fiber.ErrNotFound
	}

	return c.JSON(&User{
		User: authUser,
	})
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

	if storedUser, err := store.GetUserByName(user.User); err != nil {
		log.Error().Err(err).Send()
		return fiber.ErrInternalServerError
	} else if storedUser != nil {
		return fiber.ErrUnprocessableEntity
	}

	if err := store.CreateUser(&store.User{
		UniqueName:   user.User,
		PasswordHash: user.PasswordHash(),
	}); err != nil {
		log.Error().Err(err).Send()
		return fiber.ErrInternalServerError
	}

	return nil
}
