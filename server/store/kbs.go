package store

func ListKBs(conds map[string]interface{}) ([]*KB, error) {
	tx := DB()
	tx = tx.Where(conds)
	kbs := []*KB{}
	if err := tx.Find(&kbs).Error; err != nil {
		return nil, err
	}
	return kbs, nil
}

func GetKB(conds map[string]interface{}) (*KB, error) {
	kbs, err := ListKBs(conds)
	if err != nil {
		return nil, err
	}
	if len(kbs) == 0 {
		return nil, nil
	}
	return kbs[0], nil
}

func CreateKB(kb *KB) error {
	return DB().Create(kb).Error
}

func DeleteKB(conds map[string]interface{}) (int64, error) {
	tx := DB()
	tx = tx.Where(conds)
	tx = tx.Delete(&KB{})
	return tx.RowsAffected, tx.Error
}

func UpdateKB(
	conds map[string]interface{}, fields map[string]interface{},
) (int64, error) {
	tx := DB().Model(&KB{})
	tx = tx.Where(conds)
	tx = tx.Updates(fields)
	return tx.RowsAffected, tx.Error
}
