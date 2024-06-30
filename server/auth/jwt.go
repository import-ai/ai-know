package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/zerolog/log"
	"github.com/ycdzj/shuinotes/server/config"
)

func GenerateJWT(user string, exp time.Time) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": user,
		"exp":  exp.Unix(),
	})

	tokenStr, err := token.SignedString([]byte(config.JWTSecretKey()))
	if err != nil {
		log.Error().Err(err).Send()
		return "", err
	}
	return tokenStr, nil
}

func ValidateJWT(tokenStr string) (string, bool) {
	if tokenStr == "" {
		return "", false
	}

	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.JWTSecretKey()), nil
	})
	if err != nil {
		log.Error().Err(err).Send()
		return "", false
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", false
	}

	exp, err := claims.GetExpirationTime()
	if !exp.After(time.Now()) {
		return "", false
	}

	user, ok := claims["user"].(string)
	if !ok {
		return "", false
	}

	return user, true
}
