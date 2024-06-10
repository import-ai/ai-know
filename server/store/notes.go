package store

import (
	"time"

	"github.com/rs/zerolog/log"
)

type Note struct {
	ID        int64 `gorm:"primarykey"`
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func GetAllNotes() ([]*Note, error) {
	notes := []*Note{}
	if err := db.Find(&notes).Error; err != nil {
		log.Error().Err(err).Send()
		return nil, err
	}
	return notes, nil
}

func UpdateNoteContent(id int64, content string) error {
	tx := db.Model(&Note{})
	tx = tx.Where("id = ?", id)

	if err := tx.Update("content", content).Error; err != nil {
		log.Error().Err(err).Send()
		return err
	}
	return nil
}

func CreateNote(note *Note) error {
	if err := db.Create(note).Error; err != nil {
		log.Error().Err(err).Send()
		return err
	}
	return nil
}
