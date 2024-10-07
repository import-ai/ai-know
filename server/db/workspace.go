package db

import (
	"context"
	"errors"

	"github.com/import-ai/ai-know/server/sql/queries"
	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog/log"
)

func GetWorkspace(ctx context.Context, id int64) (*queries.Workspace, error) {
	log.Info().Int64("workspace_id", id).Send()
	workspace, err := newQueries().GetWorkspace(ctx, id)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		log.Err(err).Send()
	}
	return workspace, err
}

func CreateFirstWorkspace(ctx context.Context) (*queries.Workspace, error) {
	tx, err := connPool.Begin(ctx)
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
