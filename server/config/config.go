package config

import (
	"os"

	"github.com/rs/zerolog/log"
)

var dsn string
var listenAddr string

func InitFromEnv() {
	listenAddr = os.Getenv("API_LISTEN_ADDR")
	dsn = os.Getenv("API_DSN")
	checkConfig()
}

func checkConfig() {
	if listenAddr == "" {
		log.Fatal().Msg("Listen address is empty")
	}
	if dsn == "" {
		log.Fatal().Msg("DSN is empty")
	}
}

func DSN() string {
	return dsn
}

func ListenAddr() string {
	return listenAddr
}
