package handlers

import (
	"bytes"
	"crypto/sha256"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"github.com/ycdzj/shuinotes/server/auth"
	"github.com/ycdzj/shuinotes/server/config"
	"github.com/ycdzj/shuinotes/server/store"
)

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (s *User) PasswordHash() []byte {
	hashBytes := sha256.Sum256([]byte(s.Name + "#" + s.Password))
	return hashBytes[:]
}

func HandleLogin(c *fiber.Ctx) error {
	req := &User{}
	if err := c.BodyParser(req); err != nil {
		log.Error().Err(err).Send()
		return fiber.ErrBadRequest
	}
	if req.Name == "" || req.Password == "" {
		return fiber.ErrBadRequest
	}

	storedUser, err := store.GetUserByName(req.Name)
	if err != nil {
		log.Error().Err(err).Send()
		return fiber.ErrInternalServerError
	}
	if storedUser == nil {
		return fiber.ErrUnprocessableEntity
	}
	if !bytes.Equal(storedUser.PasswordHash, req.PasswordHash()) {
		return fiber.ErrUnprocessableEntity
	}

	expireTime := time.Now().Add(time.Hour * 24)
	tokenStr, err := auth.GenerateJWT(req.Name, expireTime)
	if err != nil {
		return fiber.ErrInternalServerError
	}
	c.Cookie(&fiber.Cookie{
		Name:    config.JWTCookieName(),
		Value:   tokenStr,
		Expires: expireTime,
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
		Name: authUser,
	})
}

func HandleCreateUser(c *fiber.Ctx) error {
	user := &User{}
	if err := c.BodyParser(user); err != nil {
		log.Error().Err(err).Send()
		return fiber.ErrBadRequest
	}
	if user.Name == "" || user.Password == "" {
		return fiber.ErrBadRequest
	}

	if storedUser, err := store.GetUserByName(user.Name); err != nil {
		log.Error().Err(err).Send()
		return fiber.ErrInternalServerError
	} else if storedUser != nil {
		return fiber.ErrUnprocessableEntity
	}

	if err := store.CreateUser(&store.User{
		Name:         user.Name,
		PasswordHash: user.PasswordHash(),
	}); err != nil {
		log.Error().Err(err).Send()
		return fiber.ErrInternalServerError
	}

	return nil
}

func getAuthorizedUser(c *fiber.Ctx) (*store.User, error) {
	userName, ok := c.Locals("user").(string)
	if !ok || userName == "" {
		return nil, nil
	}
	user, err := store.GetUserByName(userName)
	if err != nil {
		return nil, err
	}
	return user, nil
}
