package db

import (
	"context"
	"errors"

	"github.com/import-ai/ai-know/server/sql/queries"
	"github.com/import-ai/ai-know/server/sql/sqlstates"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
)

var ErrEntryNotExist = errors.New("Entry not exist")
var ErrInvalidEntry = errors.New("Invalid entry")

type PutSidebarEntryArgs struct {
	EntryID int64
	Title   string
	Parent  int64
	PrevID  int64
}

type CreateSidebarEntryArgs struct {
	Title  string
	Type   queries.SidebarEntryType
	Parent int64
	PrevID int64
}

func newValidInt8(val int64) pgtype.Int8 {
	return pgtype.Int8{
		Int64: val,
		Valid: true,
	}
}

func insertSidebarSubEntry(
	ctx context.Context, q *queries.Queries,
	entryID int64, parentID int64, prevID int64,
) error {
	if prevID != parentID {
		prevEntry, err := q.LockSidebarEntry(ctx, prevID)
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrEntryNotExist
		}
		if err != nil {
			return err
		}
		if !prevEntry.ParentID.Valid ||
			prevEntry.ParentID.Int64 != parentID {
			return ErrInvalidEntry
		}
	}

	if _, err := q.ReplacePrevEntry(ctx, &queries.ReplacePrevEntryParams{
		ParentID: newValidInt8(parentID),
		OldPrev:  newValidInt8(prevID),
		NewPrev:  newValidInt8(entryID),
	}); err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return err
	}

	_, err := q.SetParentPrevEntry(ctx, &queries.SetParentPrevEntryParams{
		ID:       entryID,
		ParentID: newValidInt8(parentID),
		PrevID:   newValidInt8(prevID),
	})
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.SQLState() == sqlstates.ForeignKeyViolation {
				return ErrEntryNotExist
			}
		}
		return err
	}
	return nil
}

func removeSidebarSubEntry(
	ctx context.Context, q *queries.Queries, entry *queries.SidebarEntry,
) error {
	if !entry.ParentID.Valid || !entry.PrevID.Valid {
		return ErrInvalidEntry
	}

	if _, err := q.SetParentPrevEntry(ctx, &queries.SetParentPrevEntryParams{
		ID: entry.ID,
	}); err != nil {
		return err
	}

	if _, err := q.ReplacePrevEntry(ctx, &queries.ReplacePrevEntryParams{
		ParentID: entry.ParentID,
		OldPrev:  newValidInt8(entry.ID),
		NewPrev:  entry.PrevID,
	}); err != nil && !errors.Is(err, pgx.ErrNoRows) {
		return err
	}

	return nil
}

func CreateSidebarEntry(
	ctx context.Context, args *CreateSidebarEntryArgs,
) (*queries.SidebarEntry, error) {
	tx, err := connPool.Begin(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)
	q := newQueries().WithTx(tx)

	entry, err := q.CreateSidebarEntry(ctx, &queries.CreateSidebarEntryParams{
		Type:  args.Type,
		Title: args.Title,
	})
	if err != nil {
		return nil, err
	}

	if err := insertSidebarSubEntry(
		ctx, q, entry.ID, args.Parent, args.PrevID,
	); err != nil {
		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}
	return entry, nil
}

func GetSidebarEntry(
	ctx context.Context, id int64,
) (*queries.SidebarEntry, error) {
	entry, err := newQueries().GetSidebarEntry(ctx, id)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	return entry, err
}

func GetSidebarSubEntry(
	ctx context.Context, parentID int64, prevID int64,
) (*queries.SidebarEntry, error) {
	entry, err := newQueries().GetSidebarSubEntry(
		ctx, &queries.GetSidebarSubEntryParams{
			ParentID: newValidInt8(parentID),
			PrevID:   newValidInt8(prevID),
		},
	)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	return entry, err
}

func lockAllParents(
	ctx context.Context, q *queries.Queries, parentID int64, entryID int64,
) error {
	for {
		if parentID == entryID {
			return ErrInvalidEntry
		}
		parent, err := q.LockSidebarEntry(ctx, parentID)
		if err != nil {
			return err
		}
		if !parent.ParentID.Valid {
			break
		}
		parentID = parent.ParentID.Int64
	}
	return nil
}

func PutSidebarEntry(
	ctx context.Context, args *PutSidebarEntryArgs,
) error {
	tx, err := connPool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)
	q := newQueries().WithTx(tx)

	if args.Title != "" {
		if _, err := q.SetEntryTitle(ctx, &queries.SetEntryTitleParams{
			ID:    args.EntryID,
			Title: args.Title,
		}); errors.Is(err, pgx.ErrNoRows) {
			return ErrEntryNotExist
		} else if err != nil {
			return err
		}
	}

	if args.Parent != 0 {
		entry, err := q.GetSidebarEntry(ctx, args.EntryID)
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrEntryNotExist
		}
		if err != nil {
			return err
		}

		if err := lockAllParents(
			ctx, q, args.Parent, args.EntryID,
		); err != nil {
			return err
		}

		if err := removeSidebarSubEntry(ctx, q, entry); err != nil {
			return err
		}

		if err := insertSidebarSubEntry(
			ctx, q, args.EntryID, args.Parent, args.PrevID,
		); err != nil {
			return err
		}
	}

	return tx.Commit(ctx)
}

func RemoveSidebarEntry(ctx context.Context, id int64) error {
	tx, err := connPool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)
	q := newQueries().WithTx(tx)

	entry, err := q.GetSidebarEntry(ctx, id)
	if errors.Is(err, pgx.ErrNoRows) {
		return ErrEntryNotExist
	}
	if err != nil {
		return err
	}
	if err := removeSidebarSubEntry(ctx, q, entry); err != nil {
		return err
	}

	return tx.Commit(ctx)
}
