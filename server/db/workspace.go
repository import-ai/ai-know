package db

import (
	"context"
	"errors"

	"github.com/import-ai/ai-know/server/sql/queries"
	"github.com/jackc/pgx/v5"
)

func GetWorkspace(ctx context.Context, id int64) (*queries.Workspace, error) {
	workspace, err := newQueries().GetWorkspace(ctx, id)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	return workspace, err
}

func CreateFirstWorkspace(ctx context.Context) (*queries.Workspace, error) {
	tx, err := conn.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)
	q := newQueries().WithTx(tx)

	privateEntry, err := q.CreateSidebarEntry(
		ctx, &queries.CreateSidebarEntryParams{
			Type:  queries.SidebarEntryTypeGroup,
			Title: "Private",
		},
	)
	if err != nil {
		return nil, err
	}

	teamEntry, err := q.CreateSidebarEntry(
		ctx, &queries.CreateSidebarEntryParams{
			Type:  queries.SidebarEntryTypeGroup,
			Title: "Team",
		},
	)
	if err != nil {
		return nil, err
	}

	workspace, err := q.CreateWorkspace(ctx, &queries.CreateWorkspaceParams{
		ID:                  1,
		PrivateSidebarEntry: privateEntry.ID,
		TeamSidebarEntry:    teamEntry.ID,
	})
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}
	return workspace, nil
}
