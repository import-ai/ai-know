// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: queries.sql

package queries

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createSidebarEntry = `-- name: CreateSidebarEntry :one
INSERT INTO sidebar_entries(type, title, parent_id, prev_id)
VALUES ($1, $2, $3, $4)
RETURNING id, type, title, parent_id, prev_id, created_at, updated_at
`

type CreateSidebarEntryParams struct {
	Type     SidebarEntryType
	Title    string
	ParentID pgtype.Int8
	PrevID   pgtype.Int8
}

func (q *Queries) CreateSidebarEntry(ctx context.Context, arg *CreateSidebarEntryParams) (*SidebarEntry, error) {
	row := q.db.QueryRow(ctx, createSidebarEntry,
		arg.Type,
		arg.Title,
		arg.ParentID,
		arg.PrevID,
	)
	var i SidebarEntry
	err := row.Scan(
		&i.ID,
		&i.Type,
		&i.Title,
		&i.ParentID,
		&i.PrevID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}

const createWorkspace = `-- name: CreateWorkspace :one
INSERT INTO workspaces(id, private_sidebar_entry, team_sidebar_entry)
VALUES ($1, $2, $3)
RETURNING id, private_sidebar_entry, team_sidebar_entry, created_at, updated_at
`

type CreateWorkspaceParams struct {
	ID                  int64
	PrivateSidebarEntry int64
	TeamSidebarEntry    int64
}

func (q *Queries) CreateWorkspace(ctx context.Context, arg *CreateWorkspaceParams) (*Workspace, error) {
	row := q.db.QueryRow(ctx, createWorkspace, arg.ID, arg.PrivateSidebarEntry, arg.TeamSidebarEntry)
	var i Workspace
	err := row.Scan(
		&i.ID,
		&i.PrivateSidebarEntry,
		&i.TeamSidebarEntry,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}

const getNextSidebarEntry = `-- name: GetNextSidebarEntry :one
SELECT id, type, title, parent_id, prev_id, created_at, updated_at
FROM sidebar_entries
WHERE parent_id = $1
  AND prev_id = $2
`

type GetNextSidebarEntryParams struct {
	ParentID pgtype.Int8
	PrevID   pgtype.Int8
}

func (q *Queries) GetNextSidebarEntry(ctx context.Context, arg *GetNextSidebarEntryParams) (*SidebarEntry, error) {
	row := q.db.QueryRow(ctx, getNextSidebarEntry, arg.ParentID, arg.PrevID)
	var i SidebarEntry
	err := row.Scan(
		&i.ID,
		&i.Type,
		&i.Title,
		&i.ParentID,
		&i.PrevID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}

const getSidebarEntry = `-- name: GetSidebarEntry :one
SELECT id, type, title, parent_id, prev_id, created_at, updated_at
FROM sidebar_entries
WHERE id = $1
`

func (q *Queries) GetSidebarEntry(ctx context.Context, id int64) (*SidebarEntry, error) {
	row := q.db.QueryRow(ctx, getSidebarEntry, id)
	var i SidebarEntry
	err := row.Scan(
		&i.ID,
		&i.Type,
		&i.Title,
		&i.ParentID,
		&i.PrevID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}

const getWorkspace = `-- name: GetWorkspace :one
SELECT id, private_sidebar_entry, team_sidebar_entry, created_at, updated_at
FROM workspaces
WHERE id = $1
`

func (q *Queries) GetWorkspace(ctx context.Context, id int64) (*Workspace, error) {
	row := q.db.QueryRow(ctx, getWorkspace, id)
	var i Workspace
	err := row.Scan(
		&i.ID,
		&i.PrivateSidebarEntry,
		&i.TeamSidebarEntry,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}

const lockSidebarEntry = `-- name: LockSidebarEntry :one
SELECT id, type, title, parent_id, prev_id, created_at, updated_at
FROM sidebar_entries
WHERE id = $1 FOR UPDATE
`

func (q *Queries) LockSidebarEntry(ctx context.Context, id int64) (*SidebarEntry, error) {
	row := q.db.QueryRow(ctx, lockSidebarEntry, id)
	var i SidebarEntry
	err := row.Scan(
		&i.ID,
		&i.Type,
		&i.Title,
		&i.ParentID,
		&i.PrevID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}

const replacePrevEntry = `-- name: ReplacePrevEntry :one
UPDATE sidebar_entries
SET prev_id = $1
WHERE parent_id = $2
  AND prev_id = $3
RETURNING id, type, title, parent_id, prev_id, created_at, updated_at
`

type ReplacePrevEntryParams struct {
	NewPrev  pgtype.Int8
	ParentID pgtype.Int8
	OldPrev  pgtype.Int8
}

func (q *Queries) ReplacePrevEntry(ctx context.Context, arg *ReplacePrevEntryParams) (*SidebarEntry, error) {
	row := q.db.QueryRow(ctx, replacePrevEntry, arg.NewPrev, arg.ParentID, arg.OldPrev)
	var i SidebarEntry
	err := row.Scan(
		&i.ID,
		&i.Type,
		&i.Title,
		&i.ParentID,
		&i.PrevID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}

const setPrevEntry = `-- name: SetPrevEntry :one
UPDATE sidebar_entries
SET prev_id = $2
WHERE id = $1
  AND prev_id IS NULL
RETURNING id, type, title, parent_id, prev_id, created_at, updated_at
`

type SetPrevEntryParams struct {
	ID     int64
	PrevID pgtype.Int8
}

func (q *Queries) SetPrevEntry(ctx context.Context, arg *SetPrevEntryParams) (*SidebarEntry, error) {
	row := q.db.QueryRow(ctx, setPrevEntry, arg.ID, arg.PrevID)
	var i SidebarEntry
	err := row.Scan(
		&i.ID,
		&i.Type,
		&i.Title,
		&i.ParentID,
		&i.PrevID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}
