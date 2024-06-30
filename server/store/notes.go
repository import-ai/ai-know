package store

func ListNotes(conds map[string]interface{}) ([]*Note, error) {
	tx := DB()
	tx = tx.Where(conds)
	notes := []*Note{}
	if err := tx.Find(&notes).Error; err != nil {
		return nil, err
	}
	return notes, nil
}

func GetNote(conds map[string]interface{}) (*Note, error) {
	notes, err := ListNotes(conds)
	if err != nil {
		return nil, err
	}
	if len(notes) == 0 {
		return nil, nil
	}
	return notes[0], nil
}

func CreateNote(note *Note) error {
	return DB().Create(note).Error
}

func DeleteNote(conds map[string]interface{}) (int64, error) {
	tx := DB()
	tx = tx.Where(conds)
	tx = tx.Delete(&Note{})
	return tx.RowsAffected, tx.Error
}

func UpdateNote(
	conds map[string]interface{}, fields map[string]interface{},
) (int64, error) {
	tx := DB().Model(&Note{})
	tx = tx.Where(conds)
	tx = tx.Updates(fields)
	return tx.RowsAffected, tx.Error
}
