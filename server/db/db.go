package db

import (
	"context"

	"github.com/import-ai/ai-know/server/config"
	"github.com/import-ai/ai-know/server/sql/queries"
	"github.com/jackc/pgx/v5"
	"github.com/rs/zerolog/log"
)

var conn *pgx.Conn

func Init(ctx context.Context) error {
	dsn := config.DataSourceName()
	if dsn == "" {
		log.Fatal().Msg("Data Source Name is empty")
	}
	var err error
	conn, err = pgx.Connect(ctx, dsn)
	return err
}

func Close(ctx context.Context) error {
	return conn.Close(ctx)
}

func newQueries() *queries.Queries {
	return queries.New(conn)
}
