package handlers

import (
	"bytes"
	"crypto/sha256"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/import-ai/ai-know/server/auth"
	"github.com/import-ai/ai-know/server/config"
	"github.com/import-ai/ai-know/server/store"
	"github.com/import-ai/ai-know/server/utils"
	"github.com/rs/zerolog/log"
)

type User struct {
	Name     string `json:"name,omitempty"`
	Password string `json:"password,omitempty"`
}

func (s *User) PasswordHash() []byte {
	hashBytes := sha256.Sum256([]byte(s.Name + "#" + s.Password))
	return hashBytes[:]
}

func parseUserReq(c *fiber.Ctx) (*User, error) {
	req := &User{}
	if err := c.BodyParser(req); err != nil {
		log.Error().Err(err).Send()
		return nil, utils.MakeErrorResp(c, fiber.StatusBadRequest, "Not a valid json")
	}
	if req.Name == "" || req.Password == "" {
		return nil, utils.MakeErrorResp(c, fiber.StatusBadRequest, "Empty name or password")
	}
	return req, nil
}

func HandleLogin(c *fiber.Ctx) error {
	req, err := parseUserReq(c)
	if req == nil {
		return err
	}

	storedUser, err := store.GetUserByName(req.Name)
	if err != nil {
		log.Error().Err(err).Send()
		return utils.MakeErrorResp(c, fiber.StatusInternalServerError, "")
	}
	if storedUser == nil || !bytes.Equal(storedUser.PasswordHash, req.PasswordHash()) {
		return utils.MakeErrorResp(c, fiber.StatusUnprocessableEntity, "Incorrect username or password")
	}

	expireTime := time.Now().Add(time.Hour * 24)
	tokenStr, err := auth.GenerateJWT(req.Name, expireTime)
	if err != nil {
		return utils.MakeErrorResp(c, fiber.StatusInternalServerError, "")
	}
	c.Cookie(&fiber.Cookie{
		Name:    config.JWTCookieName(),
		Value:   tokenStr,
		Expires: expireTime,
	})
	return utils.MakeOKResp(c, &User{
		Name: req.Name,
	})
}

func HandleLogout(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:  config.JWTCookieName(),
		Value: "",
	})
	return utils.MakeOKResp(c, &User{})
}

func HandleGetAuthorizedUser(c *fiber.Ctx) error {
	user, ok := c.Locals("authorized_user").(*store.User)
	if !ok {
		return utils.MakeErrorResp(c, fiber.StatusInternalServerError, "")
	}

	return utils.MakeOKResp(c, &User{
		Name: user.Name,
	})
}

func HandleRegister(c *fiber.Ctx) error {
	req, err := parseUserReq(c)
	if req == nil {
		return err
	}

	if storedUser, err := store.GetUserByName(req.Name); err != nil {
		log.Error().Err(err).Send()
		return utils.MakeErrorResp(c, fiber.StatusInternalServerError, "")
	} else if storedUser != nil {
		return utils.MakeErrorResp(c, fiber.StatusUnprocessableEntity, "User already exists")
	}

	if err := store.CreateUser(&store.User{
		Name:         req.Name,
		PasswordHash: req.PasswordHash(),
	}); err != nil {
		log.Error().Err(err).Send()
		return utils.MakeErrorResp(c, fiber.StatusInternalServerError, "")
	}

	return utils.MakeOKResp(c, &User{
		Name: req.Name,
	})
}
