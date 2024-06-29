package handlers

import (
	"crypto/sha256"
	"encoding/hex"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
	"github.com/ycdzj/shuinotes/server/config"
	"github.com/ycdzj/shuinotes/server/store"
)

type LoginReq struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

func getPasswordHash(password string) string {
	hashBytes := sha256.Sum256([]byte(password))
	return hex.EncodeToString(hashBytes[:])
}

func HandleLogin(c *fiber.Ctx) error {
	req := &LoginReq{}
	if err := c.BodyParser(req); err != nil {
		log.Error().Err(err).Send()
		return fiber.ErrBadRequest
	}
	if req.User == "" || req.Password == "" {
		return fiber.ErrBadRequest
	}

	user, err := store.GetUserByName(req.User)
	if err != nil {
		log.Error().Err(err).Send()
		return fiber.ErrInternalServerError
	}

	hash := getPasswordHash(req.Password)
	if user == nil || user.PasswordHash != hash {
		return fiber.ErrUnprocessableEntity
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": user.UniqueName,
		"exp":  time.Now().Add(time.Hour * 24),
	})

	tokenStr, err := token.SignedString(config.JWTSecretKey())
	if err != nil {
		log.Error().Err(err).Send()
		return fiber.ErrInternalServerError
	}

	c.Cookie(&fiber.Cookie{
		Name:  config.JWTCookieName(),
		Value: tokenStr,
	})

	return c.SendStatus(fiber.StatusOK)
}
