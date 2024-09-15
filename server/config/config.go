package config

import (
	"os"
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
