package store

import "gorm.io/gorm"

func getFirstUser(tx *gorm.DB) (*User, error) {
	user := &User{}
	if err := tx.First(user).Error; err == gorm.ErrRecordNotFound {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return user, nil
}

func GetUserByName(uniqueName string) (*User, error) {
	tx := DB()
	tx = tx.Where("unique_name = ?", uniqueName)
	return getFirstUser(tx)
}

func GetUserByID(id int64) (*User, error) {
	tx := DB()
	tx = tx.Where("id = ?", id)
	return getFirstUser(tx)
}

func CreateUser(user *User) error {
	return DB().Create(user).Error
}
