package stores

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type PostgresStore struct {
	ctx context.Context
	db  *pgx.Conn
}

func NewPostgresStore(
	ctx context.Context,
	db *pgx.Conn,
) *PostgresStore {
	return &PostgresStore{
		ctx: ctx,
		db:  db,
	}
}

func (s *PostgresStore) Ctx() context.Context {
	return s.ctx
}

func (s *PostgresStore) DB() *pgx.Conn {
	return s.db
}
