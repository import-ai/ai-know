package store

import "time"

type User struct {
	ID           int64  `gorm:"primarykey;autoIncrement"`
	UniqueName   string `gorm:"uniqueIndex"` // The unique user name used for login
	PasswordHash string // SHA256 sum of the user's password
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// Knowledge Base
type KB struct {
	ID          int64  `gorm:"primarykey;autoIncrement"`
	UniqueName  string `gorm:"uniqueIndex"` // The unique identifier that may be exposed externally
	OwnerUserID int64  `gorm:"index"`
	Title       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Note struct {
	ID           int64  `gorm:"primarykey;autoIncrement"`
	UniqueName   string `gorm:"uniqueIndex"` // The unique identifier that may be exposed externally
	KBID         int64  `gorm:"index:idx_kb_par,priority:1"`
	ParentNoteID int64  `gorm:"index:idx_kb_par,priority:2"` // Parent must belong to the same KB
	Title        string
	Content      string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
