package store

import (
	"time"

	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Note struct {
	ID        int64 `gorm:"primarykey"`
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func GetAllNotes() ([]*Note, error) {
	tx := db.Model(&Note{})
	tx = tx.Order(clause.OrderByColumn{
		Column: clause.Column{Name: "id"},
		Desc:   true,
	})

	notes := []*Note{}
	if err := tx.Find(&notes).Error; err != nil {
		log.Error().Err(err).Send()
		return nil, err
	}
	return notes, nil
}

func GetNoteByID(id int64) (*Note, error) {
	tx := db.Model(&Note{})
	tx = tx.Where("id = ?", id)

	note := &Note{}
	if err := tx.First(note).Error; err == gorm.ErrRecordNotFound {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return note, nil
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
