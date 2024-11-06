package db

import (
	"context"
	"errors"

	"github.com/import-ai/ai-know/server/config"
	"github.com/import-ai/ai-know/server/sql/queries"
	"github.com/jackc/pgx/v5/pgxpool"
)

var connPool *pgxpool.Pool

func Init(ctx context.Context) error {
	dsn := config.DataSourceName()
	if dsn == "" {
		return errors.New("Data Source Name is empty")
	}
	var err error
	connPool, err = pgxpool.New(ctx, dsn)
	return err
}

func Close(ctx context.Context) error {
	connPool.Close()
	return nil
}

func newQueries() *queries.Queries {
	return queries.New(connPool)
}
