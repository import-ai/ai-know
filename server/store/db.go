package store

import (
	"os"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() error {
	dsn := os.Getenv("DB_DSN")
	// dsn := "host=localhost user=gorm password=gorm dbname=gorm port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Error().Err(err).Send()
		return err
	}
	db = db.Debug()

	log.Info().Msg("DB initialized")
	return nil
}

func AutoMigrate() error {
	if err := db.AutoMigrate(&Note{}); err != nil {
		log.Error().Err(err).Send()
		return err
	}
	return nil
}
