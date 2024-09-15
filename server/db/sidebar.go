package db

import (
	"context"
	"errors"

	"github.com/import-ai/ai-know/server/sql/queries"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

var ErrEntryNotExist = errors.New("Entry not exist")
var ErrInvalidPosition = errors.New("Invalid position")

type CreateSidebarEntryArgs struct {
	Title         string
	Type          queries.SidebarEntryType
	Parent        int64
	PositionAfter int64
}

func newValidInt8(val int64) pgtype.Int8 {
	return pgtype.Int8{
		Int64: val,
		Valid: true,
	}
}

func CreateSidebarEntry(
	ctx context.Context, args *CreateSidebarEntryArgs,
) (*queries.SidebarEntry, error) {
	tx, err := conn.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)
	q := newQueries().WithTx(tx)

	_, err = q.GetSidebarEntry(ctx, args.Parent)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, ErrEntryNotExist
	}
	if err != nil {
		return nil, err
	}

	if args.PositionAfter == 0 {
		args.PositionAfter = args.Parent
	} else {
		prevEntry, err := q.LockSidebarEntry(ctx, args.PositionAfter)
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrEntryNotExist
		}
		if err != nil {
			return nil, err
		}
		if !prevEntry.ParentID.Valid ||
			prevEntry.ParentID.Int64 != args.Parent {
			return nil, ErrInvalidPosition
		}
	}

	entry, err := q.CreateSidebarEntry(ctx, &queries.CreateSidebarEntryParams{
		Type:     args.Type,
		Title:    args.Title,
		ParentID: newValidInt8(args.Parent),
	})
	if err != nil {
		return nil, err
	}

	if _, err := q.ReplacePrevEntry(ctx, &queries.ReplacePrevEntryParams{
		NewPrev:  newValidInt8(entry.ID),
		ParentID: newValidInt8(args.Parent),
		OldPrev:  newValidInt8(args.PositionAfter),
	}); err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return nil, err
	}

	entry, err = q.SetPrevEntry(ctx, &queries.SetPrevEntryParams{
		ID:     entry.ID,
		PrevID: newValidInt8(args.PositionAfter),
	})
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}
	return entry, nil
}
