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
	"github.com/ycdzj/shuinotes/server/utils"
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
	if storedUser == nil {
		return utils.MakeErrorResp(c, fiber.StatusUnprocessableEntity, "User not exist")
	}
	if !bytes.Equal(storedUser.PasswordHash, req.PasswordHash()) {
		return utils.MakeErrorResp(c, fiber.StatusUnprocessableEntity, "Incorrect password")
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

func HandleGetUser(c *fiber.Ctx) error {
	authUser, ok := c.Locals("user").(string)
	if !ok {
		return utils.MakeErrorResp(c, fiber.StatusInternalServerError, "")
	}

	if c.Params("user_name") != authUser {
		return utils.MakeErrorResp(c, fiber.StatusNotFound, "Not found")
	}

	return utils.MakeOKResp(c, &User{
		Name: authUser,
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
		return fiber.ErrUnprocessableEntity
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

func getAuthorizedUser(c *fiber.Ctx) (*store.User, error) {
	userName, ok := c.Locals("user").(string)
	if !ok || userName == "" {
		return nil, utils.MakeErrorResp(c, fiber.StatusUnauthorized, "Not logged in")
	}
	user, err := store.GetUserByName(userName)
	if err != nil {
		log.Error().Err(err).Send()
		return nil, utils.MakeErrorResp(c, fiber.StatusInternalServerError, "")
	}
	if user == nil {
		return nil, utils.MakeErrorResp(c, fiber.StatusUnauthorized, "Not logged in")
	}
	return user, nil
}
