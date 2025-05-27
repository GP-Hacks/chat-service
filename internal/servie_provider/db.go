package servie_provider

import (
	"context"

	"github.com/GP-Hacks/chat/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func (s *ServiceProvider) DB() *pgxpool.Pool {
	if s.db == nil {
		pool, err := pgxpool.New(context.Background(), "postgres://"+config.Cfg.Postgres.User+":"+config.Cfg.Postgres.Password+"@"+config.Cfg.Postgres.Address+"/"+config.Cfg.Postgres.Name+"?sslmode=disable")
		if err != nil {
			// log.Fatal().Msg("Failed connection to PostgreSQL")
		}

		s.db = pool
	}

	return s.db
}
