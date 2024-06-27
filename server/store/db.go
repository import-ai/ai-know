package store

import (
	"github.com/rs/zerolog/log"
	"github.com/ycdzj/shuinotes/server/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var allModels []interface{}
var db *gorm.DB

func InitDB() error {
	var err error
	db, err = gorm.Open(postgres.Open(config.DSN()), &gorm.Config{})
	if err != nil {
		log.Error().Err(err).Send()
		return err
	}
	db = db.Debug()

	log.Info().Msg("DB initialized")
	return nil
}

func AutoMigrate() error {
	if err := db.AutoMigrate(allModels...); err != nil {
		log.Error().Err(err).Send()
		return err
	}
	return nil
}
