package store

import (
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var allModels []interface{}
var db *gorm.DB

func InitDB(dataSourceName string) error {
	var err error
	db, err = gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		log.Error().Err(err).Send()
		return err
	}

	log.Info().Msg("DB initialized")
	return nil
}

func DB() *gorm.DB {
	return db.Debug()
}

func AutoMigrate() error {
	if err := DB().AutoMigrate(allModels...); err != nil {
		log.Error().Err(err).Send()
		return err
	}
	return nil
}
