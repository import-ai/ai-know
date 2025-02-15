// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package queries

import (
	"database/sql/driver"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
)

type SidebarEntryType string

const (
	SidebarEntryTypeGroup SidebarEntryType = "group"
	SidebarEntryTypeNote  SidebarEntryType = "note"
	SidebarEntryTypeLink  SidebarEntryType = "link"
)

func (e *SidebarEntryType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = SidebarEntryType(s)
	case string:
		*e = SidebarEntryType(s)
	default:
		return fmt.Errorf("unsupported scan type for SidebarEntryType: %T", src)
	}
	return nil
}

type NullSidebarEntryType struct {
	SidebarEntryType SidebarEntryType
	Valid            bool // Valid is true if SidebarEntryType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullSidebarEntryType) Scan(value interface{}) error {
	if value == nil {
		ns.SidebarEntryType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.SidebarEntryType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullSidebarEntryType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.SidebarEntryType), nil
}

type SidebarEntry struct {
	ID        int64
	Type      SidebarEntryType
	Title     string
	ParentID  pgtype.Int8
	PrevID    pgtype.Int8
	CreatedAt pgtype.Timestamptz
	UpdatedAt pgtype.Timestamptz
}

type Workspace struct {
	ID                  int64
	PrivateSidebarEntry int64
	TeamSidebarEntry    int64
	CreatedAt           pgtype.Timestamptz
	UpdatedAt           pgtype.Timestamptz
}
