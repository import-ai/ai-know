package store

import "time"

type User struct {
	ID           int64     `gorm:"primarykey;autoIncrement"`
	Name         string    `gorm:"not null;uniqueIndex"` // The unique user name used for login
	PasswordHash []byte    `gorm:"not null"`             // SHA256 sum of the user's password
	CreatedAt    time.Time `gorm:"not null"`
	UpdatedAt    time.Time `gorm:"not null"`
}

// Knowledge Base
type KB struct {
	ID          int64     `gorm:"primarykey;autoIncrement"`
	ExternalID  string    `gorm:"not null;uniqueIndex"` // The unique identifier that may be exposed externally
	OwnerUserID int64     `gorm:"not null;index"`
	Title       string    `gorm:"not null"`
	CreatedAt   time.Time `gorm:"not null"`
	UpdatedAt   time.Time `gorm:"not null"`
}

type Note struct {
	ID           int64     `gorm:"primarykey;autoIncrement"`
	ExternalID   string    `gorm:"not null;uniqueIndex"` // The unique identifier that may be exposed externally
	KBID         int64     `gorm:"not null;index:idx_kb_par,priority:1"`
	ParentNoteID int64     `gorm:"not null;index:idx_kb_par,priority:2"` // Parent must belong to the same KB
	Title        string    `gorm:"not null"`
	Content      string    `gorm:"not null"`
	CreatedAt    time.Time `gorm:"not null"`
	UpdatedAt    time.Time `gorm:"not null"`
}

func init() {
	allModels = append(allModels, &User{})
	allModels = append(allModels, &KB{})
	allModels = append(allModels, &Note{})
}
