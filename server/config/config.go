package config

import (
	"os"

	"github.com/rs/zerolog/log"
)

var dataSourceName string
var listenAddr string
var jwtSecretKey string
var aiServerAddr string

func InitFromEnv() {
	listenAddr = os.Getenv("API_LISTEN_ADDR")
	dataSourceName = os.Getenv("API_DATA_SOURCE_NAME")
	jwtSecretKey = os.Getenv("API_JWT_SECRET_KEY")
	aiServerAddr = os.Getenv("API_AI_SERVER_ADDR")
	checkConfig()
}

func checkConfig() {
	if listenAddr == "" {
		log.Fatal().Msg("Listen Address is empty")
	}
	if dataSourceName == "" {
		log.Fatal().Msg("Data Source Name is empty")
	}
	if jwtSecretKey == "" {
		log.Fatal().Msg("JWT Secret Key is empty")
	}
	if aiServerAddr == "" {
		log.Fatal().Msg("AI Server Address is empty")
	}
}

func DataSourceName() string {
	return dataSourceName
}

func ListenAddr() string {
	return listenAddr
}

func JWTSecretKey() string {
	return jwtSecretKey
}

func AIServerAddr() string {
	return aiServerAddr
}

func JWTCookieName() string {
	return "token"
}
